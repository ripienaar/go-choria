// Copyright (c) 2022-2026, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package ccmregistrationwatcher

import (
	"encoding/json"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	cmodel "github.com/choria-io/ccm/model"
	"github.com/choria-io/go-choria/aagent/watchers/event"
)

// StateNotification describes the current state of the watcher
// described by io.choria.machine.watcher.gossip.v1.state
type StateNotification struct {
	event.Event
	Entry     cmodel.RegistrationEntry
	Published int64 `json:"previous_registration"`
}

// CloudEvent creates a CloudEvent from the state notification
func (s *StateNotification) CloudEvent() cloudevents.Event {
	return s.Event.CloudEvent(s)
}

// JSON creates a JSON representation of the notification
func (s *StateNotification) JSON() ([]byte, error) {
	return json.Marshal(s.CloudEvent())
}

// String is a string representation of the notification suitable for printing
func (s *StateNotification) String() string {
	return fmt.Sprintf("%s %s#%s cluster: %s, service: %s, port: %s, previous: %v", s.Identity, s.Machine, s.Name, s.Entry.Cluster, s.Entry.Service, s.Entry.Port, s.Published)
}
