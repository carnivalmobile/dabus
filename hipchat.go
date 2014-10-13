package main

import (
	"fmt"

	"github.com/andybons/hipchat"
)

const (
	activeMessage  = "Service %s is active"
	failedMessage  = "Service %s failed"
	restartMessage = "Service %s is auto-restarted"
)

type HipchatClient interface {
	PostMessage(req hipchat.MessageRequest) error
}

type Hipchat struct {
	Room    string `yaml:"room,omitempty"`
	Token   string `yaml:"token,omitempty"`
	Active  bool   `yaml:"on_active,omitempty"`
	Failed  bool   `yaml:"on_failed,omitempty"`
	Restart bool   `yaml:"on_restart,omitempty"`
}

func (h *Hipchat) Send(event *ServiceEvent) error {
	client := &hipchat.Client{AuthToken: h.Token}
	return h.SendWithClient(client, event)
}

func (h *Hipchat) SendWithClient(client HipchatClient, event *ServiceEvent) error {
	switch {
	case event.ActiveStatus == "active" && h.Active:
		return h.sendActive(client, event)
	case event.ActiveStatus == "failed" && h.Failed:
		return h.sendFailed(client, event)
	case event.ActiveStatus == "activating" && event.SubStatus == "auto-restart" && h.Restart:
		return h.sendRestart(client, event)
	}

	return nil
}

func (h *Hipchat) sendActive(client HipchatClient, event *ServiceEvent) error {
	msg := fmt.Sprintf(activeMessage, event.Service)
	return h.send(client, hipchat.ColorGreen, hipchat.FormatHTML, msg, false)
}

func (h *Hipchat) sendFailed(client HipchatClient, event *ServiceEvent) error {
	msg := fmt.Sprintf(failedMessage, event.Service)
	return h.send(client, hipchat.ColorRed, hipchat.FormatHTML, msg, true)
}

func (h *Hipchat) sendRestart(client HipchatClient, event *ServiceEvent) error {
	msg := fmt.Sprintf(restartMessage, event.Service)
	return h.send(client, hipchat.ColorGreen, hipchat.FormatHTML, msg, false)
}

func (h *Hipchat) send(client HipchatClient, color, format, message string, notify bool) error {
	req := hipchat.MessageRequest{
		RoomId:        h.Room,
		From:          "Systemd",
		Message:       message,
		Color:         color,
		MessageFormat: format,
		Notify:        notify,
	}

	return client.PostMessage(req)
}
