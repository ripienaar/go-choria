// Copyright (c) 2017-2021, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"syscall"
	"text/template"
	"time"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kingpin"
	"github.com/choria-io/appbuilder/builder"
	"github.com/choria-io/appbuilder/commands/exec"
	"github.com/choria-io/appbuilder/commands/parent"
	"github.com/choria-io/go-choria/protocol"
	"github.com/choria-io/go-choria/providers/appbuilder/discover"
	"github.com/choria-io/go-choria/providers/appbuilder/kv"
	"github.com/choria-io/go-choria/providers/appbuilder/rpc"
	log "github.com/sirupsen/logrus"

	"github.com/choria-io/go-choria/build"
	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-choria/config"
)

type application struct {
	app      *kingpin.Application
	command  string
	commands []runableCmd
}

var (
	cli        = application{}
	debug      = false
	configFile = ""
	c          *choria.Framework
	cfg        *config.Config
	ctx        context.Context
	cancel     func()
	wg         *sync.WaitGroup
	mu         = &sync.Mutex{}
	cpuProfile string
	bi         *build.Info
	err        error
	ran        bool
)

var usageTemplate = `{{define "FormatCommand"}}\
{{if .FlagSummary}} {{.FlagSummary}}{{end}}\
{{range .Args}} {{if not .Required}}[{{end}}<{{.Name}}>{{if .Value|IsCumulative}}...{{end}}{{if not .Required}}]{{end}}{{end}}\
{{end}}\

{{define "FormatCommands"}}\
{{range .Commands}}\
{{if not .Hidden}}\
  {{.FullCommand}}{{if .Default}}*{{end}}{{template "FormatCommand" .}}
{{.Help|Wrap 4}}
{{end}}\
{{end}}\
{{end}}\

{{ define "FormatCommandsForTopLevel" }}\
{{range .Commands}}\
{{if not .Hidden}}\
{{if not (eq .FullCommand "help")}}\
  {{.FullCommand}}{{if .Default}}*{{end}}{{template "FormatCommand" .}}
{{.Help|FirstLine|Wrap 4}}
{{end}}\
{{end}}\
{{end}}\
{{end}}\

{{define "FormatUsage"}}\
{{template "FormatCommand" .}}{{if .Commands}} <command> [<args> ...]{{end}}
{{if .Help}}
{{.Help|Wrap 0}}\
{{end}}\
{{end}}\

{{if .Context.SelectedCommand}}\
usage: {{.App.Name}} {{.Context.SelectedCommand}}{{template "FormatUsage" .Context.SelectedCommand}}
{{else}}\
usage: {{.App.Name}}{{template "FormatUsage" .App}}
{{end}}\
{{if .Context.SelectedCommand}}\
{{if .Context.Flags}}\
Flags:
{{.Context.Flags|FlagsToTwoColumns|FormatTwoColumns}}
{{end}}\
{{if .Context.Args}}\
Args:
{{.Context.Args|ArgsToTwoColumns|FormatTwoColumns}}
{{end}}\
{{if len .Context.SelectedCommand.Commands}}\
Subcommands:
{{template "FormatCommands" .Context.SelectedCommand}}
{{end}}\
{{else if .App.Commands}}\
Commands:
{{template "FormatCommandsForTopLevel" .App}}
{{end}}\
`

func ParseCLI() (err error) {
	ctx, cancel = context.WithCancel(context.Background())

	go interruptWatcher()

	// If we are not invoked as something something choria, then check
	// if the app builder has an app configuration matching the name we
	// are run as, if it does, we invoke it instead of the standard choria
	// cli tools
	//
	// TODO: too janky, need to do a better job here, looking at the name is not enough
	if !strings.Contains(os.Args[0], "choria") {
		runBuilder()
	}

	bi = &build.Info{}

	cli.app = kingpin.New("choria", "Choria Orchestration System")
	cli.app.Version(bi.Version())
	cli.app.UsageTemplate(usageTemplate)
	cli.app.UsageFuncs(template.FuncMap{
		"FirstLine": func(v string) string {
			if v == "" {
				return v
			}

			scanner := bufio.NewScanner(strings.NewReader(v))
			scanner.Scan()
			return scanner.Text()
		},
	})

	cli.app.Flag("debug", "Enable debug logging").BoolVar(&debug)
	cli.app.Flag("profile", "Enable CPU profiling and write to the supplied file").Hidden().StringVar(&cpuProfile)

	for _, cmd := range cli.commands {
		err = cmd.Setup()
	}

	cli.command = kingpin.MustParse(cli.app.Parse(os.Args[1:]))

	for _, cmd := range cli.commands {
		if cmd.FullCommand() == cli.command {
			err = cmd.Configure()
			if err != nil {
				return fmt.Errorf("%s failed to configure: %s", cmd.FullCommand(), err)
			}
		}
	}

	return
}

func builderOptions() []builder.Option {
	kv.Register()
	rpc.Register()
	discover.Register()
	parent.Register()
	exec.Register()

	return []builder.Option{builder.WithConfigPaths(
		".",
		filepath.Join(xdg.ConfigHome, "choria", "builder"),
		filepath.Join("/", "etc", "choria", "builder"),
		filepath.Join(xdg.ConfigHome, "appbuilder"),
		filepath.Join("/", "etc", "appbuilder"),
	)}
}

func runBuilder() {
	err := builder.RunStandardCLI(ctx, filepath.Base(os.Args[0]), false, nil, builderOptions()...)
	if err != nil {
		if strings.Contains(err.Error(), "must select a") {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("app builder setup failed: %v\n", err)
		}
		os.Exit(1)
	}

	os.Exit(0)
}

func systemConfigureIfRoot(actAsServer bool) error {
	if debug {
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging at debug level due to CLI override")
	}

	if configFile == "" && os.Geteuid() == 0 {
		return fmt.Errorf("configuration file must be set using --config")
	}

	if os.Geteuid() == 0 {
		cfg, err = config.NewSystemConfig(configFile, actAsServer)
	} else {
		cfg, err = config.NewConfig(configFile)
	}
	if err != nil {
		return err
	}

	applyBuildAndEnvironmentSettings()

	return nil
}

func applyBuildAndEnvironmentSettings() {
	cfg.ApplyBuildSettings(bi)

	if os.Getenv("INSECURE_ANON_TLS") == "true" {
		cfg.Choria.ClientAnonTLS = true
		cfg.DisableTLSVerify = true
		cfg.DisableSecurityProviderVerify = true
		log.Warn("Using anonymous TLS via environment override")
	}

	if os.Getenv("INSECURE_DISABLE_TLS") == "true" {
		cfg.DisableTLS = true
		log.Warn("Disabling TLS via environment override")
	}

	if os.Getenv("INSECURE_YES_REALLY") == "true" {
		cfg.DisableSecurityProviderVerify = true
		protocol.Secure = "false"
		cfg.DisableTLS = true
		log.Warn("Disabling protocol security via environment override")
	}
}

func commonConfigure() error {
	if debug {
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging at debug level due to CLI override")
	}

	if configFile == "" {
		configFile = choria.UserConfig()
	}

	cfg, err = config.NewConfig(configFile)
	if err != nil {
		return fmt.Errorf("could not parse configuration: %s", err)
	}

	applyBuildAndEnvironmentSettings()

	return nil
}

func Run() (err error) {
	wg = &sync.WaitGroup{}

	if cpuProfile != "" {
		cpf, err := os.Create(cpuProfile)
		if err != nil {
			return fmt.Errorf("could not setup profiling: %s", err)
		}
		defer cpf.Close()

		err = pprof.StartCPUProfile(cpf)
		if err != nil {
			return fmt.Errorf("could not setup profiling: %s", err)
		}

		defer pprof.StopCPUProfile()
	}

	if cfg != nil && c == nil {
		if debug {
			cfg.LogLevel = "debug"
		}

		// we do this here so that the command setup has a chance to fiddle the config for
		// things like disabling full verification of the security system during enrollment
		c, err = choria.NewWithConfig(cfg)
		if err != nil {
			return fmt.Errorf("could not initialize Choria: %s", err)
		}
	}

	for _, cmd := range cli.commands {
		if cmd.FullCommand() == cli.command {
			ran = true

			wg.Add(1)
			err = cmd.Run(wg)
		}
	}

	if !ran {
		err = fmt.Errorf("could not run the CLI: Invalid command %s", cli.command)
	}

	if err != nil {
		cancel()
	}

	wg.Wait()

	return
}

func forcequit() {
	grace := 10 * time.Second

	if cfg != nil {
		if cfg.SoftShutdownTimeout > 0 {
			grace = time.Duration(cfg.SoftShutdownTimeout) * time.Second
		}
	}

	<-time.After(grace)

	dumpGoRoutines()

	log.Errorf("Forced shutdown triggered after %v", grace)

	os.Exit(1)
}

func interruptWatcher() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case sig := <-sigs:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				go forcequit()

				log.Infof("Shutting down on %s", sig)
				cancel()

			case syscall.SIGQUIT:
				dumpGoRoutines()
			}
		case <-ctx.Done():
			return
		}
	}
}

func dumpGoRoutines() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().UnixNano()
	pid := os.Getpid()

	tdoutname := filepath.Join(os.TempDir(), fmt.Sprintf("choria-threaddump-%d-%d.txt", pid, now))
	memoutname := filepath.Join(os.TempDir(), fmt.Sprintf("choria-memoryprofile-%d-%d.mprof", pid, now))

	buf := make([]byte, 1<<20)
	stacklen := runtime.Stack(buf, true)

	err := os.WriteFile(tdoutname, buf[:stacklen], 0644)
	if err != nil {
		log.Errorf("Could not produce thread dump: %s", err)
		return
	}

	log.Warnf("Produced thread dump to %s", tdoutname)

	mf, err := os.Create(memoutname)
	if err != nil {
		log.Errorf("Could not produce memory profile: %s", err)
		return
	}
	defer mf.Close()

	err = pprof.WriteHeapProfile(mf)
	if err != nil {
		log.Errorf("Could not produce memory profile: %s", err)
		return
	}

	log.Warnf("Produced memory profile to %s", memoutname)
}

// digs in the application.commands structure looking for a entry with
// the given command string
func cmdWithFullCommand(command string) (cmd runableCmd, ok bool) {
	for _, cmd := range cli.commands {
		if cmd.FullCommand() == command {
			return cmd, true
		}
	}

	return cmd, false
}
