package main

import (
	"testing"

	"github.com/andybons/hipchat"
)

type MockHipchatClient struct {
	Request hipchat.MessageRequest
}

func (c *MockHipchatClient) PostMessage(req hipchat.MessageRequest) error {
	c.Request = req
	return nil
}

var client = &MockHipchatClient{}

var subject = &Hipchat{
	Room:    "SampleRoom",
	Token:   "foo",
	Active:  true,
	Failed:  true,
	Restart: true,
}

var event = &ServiceEvent{"foo.service", "active", ""}

func Test_SendActive(t *testing.T) {
	subject.SendWithClient(client, event)
	expected := hipchat.MessageRequest{
		RoomId:        "SampleRoom",
		From:          "Systemd",
		Message:       "Service foo.service is active",
		Color:         hipchat.ColorGreen,
		MessageFormat: hipchat.FormatHTML,
		Notify:        false,
	}

	if client.Request != expected {
		t.Errorf("Invalid hipchat payload. Expected: %v, got %v", expected, client.Request)
	}
}

func Test_SendFailed(t *testing.T) {
	event.ActiveStatus = "failed"

	subject.SendWithClient(client, event)
	expected := hipchat.MessageRequest{
		RoomId:        "SampleRoom",
		From:          "Systemd",
		Message:       "Service foo.service failed",
		Color:         hipchat.ColorRed,
		MessageFormat: hipchat.FormatHTML,
		Notify:        true,
	}

	if client.Request != expected {
		t.Errorf("Invalid hipchat payload. Expected: %v, got %v", expected, client.Request)
	}
}

func Test_SendRestart(t *testing.T) {
	event.ActiveStatus = "activating"
	event.SubStatus = "auto-restart"

	subject.SendWithClient(client, event)
	expected := hipchat.MessageRequest{
		RoomId:        "SampleRoom",
		From:          "Systemd",
		Message:       "Service foo.service is auto-restarted",
		Color:         hipchat.ColorGreen,
		MessageFormat: hipchat.FormatHTML,
		Notify:        true,
	}

	if client.Request != expected {
		t.Errorf("Invalid hipchat payload. Expected: %v, got %v", expected, client.Request)
	}
}
