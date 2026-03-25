// Copyright (c) 2026, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/choria-io/ccm/manager"
	cmodel "github.com/choria-io/ccm/model"
)

func NewCCMManager(w watcherLogger, noop bool, wd string, manifestPath string, registration string, data map[string]any, facts json.RawMessage) (cmodel.Manager, cmodel.Logger, error) {
	var fdata map[string]any
	err := json.Unmarshal(facts, &fdata)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid facts: %w", err)
	}

	log := NewCCMLogger(w)
	var opts []manager.Option
	if noop {
		w.Infof("Running in noop mode")
		opts = append(opts, manager.WithNoop())
	}
	if registration != "" {
		opts = append(opts, manager.WithRegistrationDestination(cmodel.JetStreamRegistrationDestination))
		opts = append(opts, manager.WithRegistrationStream(registration))
	}

	mgr, err := manager.NewManager(log, log, opts...)
	if err != nil {
		return nil, nil, err
	}
	mgr.SetFacts(fdata)
	mgr.SetExternalData(data)

	// try to figure out a sane root for things like source, file() etc in manifests
	if manifestPath != "" {
		if filepath.IsAbs(manifestPath) {
			wd = filepath.Dir(manifestPath)
		} else {
			abs, err := filepath.Abs(filepath.Join(wd, filepath.Dir(manifestPath)))
			if err == nil {
				wd = abs
			}
		}
	}
	mgr.SetWorkingDirectory(wd)

	return mgr, log, nil
}

type watcherLogger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

type logger struct {
	w      watcherLogger
	fields []string
}

func (l *logger) genFieldsList(args ...any) []string {
	fields := l.fields
	for i := 0; i < len(args); i += 2 {
		fields = append(fields, fmt.Sprintf("%s=%v", args[i].(string), args[i+1]))
	}

	return fields
}

func (l *logger) genFields(args ...any) string {
	return strings.Join(l.genFieldsList(args...), " ")
}

func (l *logger) Debug(msg string, args ...any) {
	l.w.Debugf("%s %s", msg, l.genFields(args...))
}

func (l *logger) Info(msg string, args ...any) {
	l.w.Infof("%s %s", msg, l.genFields(args...))
}

func (l *logger) Warn(msg string, args ...any) {
	l.w.Warnf("%s %s", msg, l.genFields(args...))
}

func (l *logger) Error(msg string, args ...any) {
	l.w.Errorf("%s %s", msg, l.genFields(args...))
}

func (l *logger) With(args ...any) cmodel.Logger {
	return &logger{
		w:      l.w,
		fields: l.genFieldsList(args...),
	}
}

func NewCCMLogger(m watcherLogger) cmodel.Logger {
	return &logger{
		w:      m,
		fields: []string{},
	}
}
