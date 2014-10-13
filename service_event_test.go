package main

import "testing"

type MockNotifier struct {
	event *ServiceEvent
}

func (n *MockNotifier) Send(event *ServiceEvent) error {
	n.event = event
	return nil
}

var serviceEvent = &ServiceEvent{"foo.service", "active", ""}
var notifier = &MockNotifier{}

func TestSendWithNotifier(t *testing.T) {
	serviceEvent.SendWithNotifier(notifier)

	if notifier.event != serviceEvent {
		t.Error("Notification failed")
	}
}
