package main

import "fmt"

const (
	slackEndpoint       = "https://%s.slack.com/services/hooks/incoming-webhook?token=%s"
	slackActiveMessage  = "Service *%s* is active"
	slackFailedMessage  = "Service *%s* failed"
	slackRestartMessage = "Service *%s* has auto-restarted"
)

type SlackAttachment struct {
	Fallback string   `json:"fallback"`
	Text     string   `json:"text"`
	Color    string   `json:"color"`
	MrkdwnIn []string `json:"mrkdwn_in"`
}

type SlackMessage struct {
	Channel     string            `json:"channel"`
	Username    string            `json:"username"`
	Attachments []SlackAttachment `json:"attachments"`
}

type Slack struct {
	Team    string `yaml:"team,omitempty"`
	Channel string `yaml:"channel,omitempty"`
	Token   string `yaml:"token,omitempty"`
	Active  bool   `yaml:"on_active,omitempty"`
	Failed  bool   `yaml:"on_failed,omitempty"`
	Restart bool   `yaml:"on_restart,omitempty"`
}

func (s *Slack) Send(event *ServiceEvent) error {
	client := NewNotifierHTTPClient()
	return s.SendWithClient(client, event)
}

func (s *Slack) SendWithClient(client HTTPClient, event *ServiceEvent) error {
	switch {
	case event.ActiveStatus == "active" && s.Active:
		return s.sendActive(client, event)
	case event.ActiveStatus == "failed" && s.Failed:
		return s.sendFailed(client, event)
	case event.ActiveStatus == "activating" && event.SubStatus == "auto-restart" && s.Restart:
		return s.sendRestart(client, event)
	}

	return nil
}

func (s *Slack) composeMessage(color string, message string) *SlackMessage {
	attachments := SlackAttachment{
		message, message, color, []string{"fallback", "text"},
	}

	return &SlackMessage{s.Channel, "Systemd", []SlackAttachment{attachments}}
}

func (s *Slack) sendActive(client HTTPClient, event *ServiceEvent) error {
	message := s.composeMessage("good",
		fmt.Sprintf(slackActiveMessage, event.Service))
	return s.send(client, message)
}

func (s *Slack) sendFailed(client HTTPClient, event *ServiceEvent) error {
	message := s.composeMessage("danger",
		fmt.Sprintf(slackFailedMessage, event.Service))
	return s.send(client, message)
}

func (s *Slack) sendRestart(client HTTPClient, event *ServiceEvent) error {
	message := s.composeMessage("warning",
		fmt.Sprintf(slackRestartMessage, event.Service))
	return s.send(client, message)
}

func (s *Slack) send(client HTTPClient, message *SlackMessage) error {
	url := fmt.Sprintf(slackEndpoint, s.Team, s.Token)
	return client.PostJSON(url, message)
}
