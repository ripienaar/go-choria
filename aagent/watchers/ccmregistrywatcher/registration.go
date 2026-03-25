// Copyright (c) 2026, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package ccmregistrationwatcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	ccmmodel "github.com/choria-io/ccm/model"
	"github.com/choria-io/go-choria/aagent/model"
	"github.com/choria-io/go-choria/aagent/util"
	"github.com/choria-io/go-choria/aagent/watchers/event"
	"github.com/choria-io/go-choria/aagent/watchers/watcher"
	iu "github.com/choria-io/go-choria/internal/util"
)

type State int

const (
	Stopped State = iota
	Running

	wtype   = "ccm_registration"
	version = "v1"
)

type properties struct {
	Stream       string
	Registration *ccmmodel.RegistrationEntry
}

type Watcher struct {
	*watcher.Watcher
	properties *properties

	name      string
	machine   model.Machine
	nc        *nats.Conn
	interval  time.Duration
	regCancel context.CancelFunc
	runCtx    context.Context
	state     State
	last      time.Time
	ccmMgr    ccmmodel.Manager
	ccmLogger ccmmodel.Logger

	terminate chan struct{}
	mu        *sync.Mutex
}

func New(machine model.Machine, name string, states []string, required []model.ForeignMachineState, failEvent string, successEvent string, interval string, ai time.Duration, properties map[string]any) (any, error) {
	var err error

	tw := &Watcher{
		name:      name,
		machine:   machine,
		terminate: make(chan struct{}),
		mu:        &sync.Mutex{},
	}

	tw.interval, err = iu.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	tw.Watcher, err = watcher.NewWatcher(name, wtype, ai, states, required, machine, failEvent, successEvent)
	if err != nil {
		return nil, err
	}

	err = tw.setProperties(properties)
	if err != nil {
		return nil, fmt.Errorf("could not set properties: %s", err)
	}

	return tw, nil
}

func (w *Watcher) stopRegistration() {
	w.mu.Lock()
	cancel := w.regCancel
	w.state = Stopped
	w.mu.Unlock()

	if cancel != nil {
		w.Infof("Stopping registration on transition to %s", w.machine.State())
		cancel()
	}
}

func (w *Watcher) startRegistration() {
	w.mu.Lock()
	cancel := w.regCancel
	ctx := w.runCtx
	w.mu.Unlock()

	if cancel != nil {
		return
	}

	go func() {
		tick := time.NewTicker(w.interval)
		gCtx, cancel := context.WithCancel(ctx)

		var err error

		w.mu.Lock()
		w.state = Running
		w.regCancel = cancel
		w.mu.Unlock()

		if err != nil {
			w.Errorf("Could not get a NATS connection to publish Registration")
		}

		stop := func() {
			w.mu.Lock()
			w.regCancel = nil
			w.state = Stopped
			tick.Stop()
			if w.ccmMgr != nil {
				w.ccmMgr.Close()
			}
			w.mu.Unlock()
		}

		publish := func() {
			if !w.ShouldWatch() {
				return
			}

			w.Infof("Registering while in state %v", w.machine.State())
			mgr, _, err := w.ccmManager()
			if err != nil {
				w.Errorf("Could not get CCM manager: %v", err)
				return
			}

			w.Debugf("Publishing Registration event")
			err = mgr.PublishRegistration(ctx, w.properties.Registration)
			if err != nil {
				w.Errorf("Could not get CCM manager: %v", err)
				return
			}

			w.mu.Lock()
			w.last = time.Now()
			w.mu.Unlock()
		}

		publish()

		for {
			select {
			case <-tick.C:
				publish()
			case <-gCtx.Done():
				stop()
				return
			case <-w.terminate:
				stop()
				return
			}
		}
	}()
}

func (w *Watcher) watch() {
	if !w.ShouldWatch() {
		w.stopRegistration()
		return
	}

	w.Infof("Starting registration timer")
	w.startRegistration()
}

func (w *Watcher) getManager() (ccmmodel.Manager, ccmmodel.Logger, error) {
	return iu.NewCCMManager(w, true, w.machine.Directory(), "", w.properties.Stream, w.machine.Data(), w.machine.Facts())
}

func (w *Watcher) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	w.mu.Lock()
	w.runCtx = ctx
	w.mu.Unlock()

	w.Infof("CCM Registration watcher starting with interval %v", w.interval)

	w.watch()

	for {
		select {
		case <-w.StateChangeC():
			w.watch()

		case <-w.terminate:
			w.Infof("Handling terminate notification")
			return
		case <-ctx.Done():
			w.Infof("Stopping on context interrupt")
			return
		}
	}
}

func (w *Watcher) setProperties(props map[string]any) error {
	if w.properties == nil {
		w.properties = &properties{}
	}

	err := util.ParseMapStructure(props, w.properties.Registration)
	if err != nil {
		return err
	}

	return w.validate()
}

func (w *Watcher) validate() error {
	err := w.properties.Registration.Validate()
	if err != nil {
		return err
	}

	if w.interval == 0 {
		w.interval = 15 * time.Second
	}

	if w.properties.Stream == "" {
		w.properties.Stream = "REGISTRATION"
	}

	return nil
}

func (w *Watcher) Delete() {
	close(w.terminate)
}

func (w *Watcher) CurrentState() any {
	w.mu.Lock()
	defer w.mu.Unlock()

	s := &StateNotification{
		Event:     event.New(w.name, wtype, version, w.machine),
		Published: w.last.Unix(),
		Entry:     *w.properties.Registration,
	}

	return s
}

func (w *Watcher) ccmManager() (ccmmodel.Manager, ccmmodel.Logger, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	var err error

	if w.ccmMgr == nil || w.ccmLogger == nil {
		w.ccmMgr, w.ccmLogger, err = iu.NewCCMManager(w, true, w.machine.Directory(), "", w.properties.Stream, w.machine.Data(), w.machine.Facts())
		if err != nil {
			return nil, nil, err
		}
	}

	return w.ccmMgr, w.ccmLogger, nil
}
