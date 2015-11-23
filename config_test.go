package main

import "testing"

const config = `
services:
  - service1
  - service2
notify:
  hipchat:
    room:  room
    token: token
    on_active:  true
    on_failed:  true
    on_restart: true
  slack:
    webhook_url: webhook_url
    channel:     channel
    username:    username
    on_active:   true
    on_failed:   true
    on_restart:  true
`

func TestNewConfigServices(t *testing.T) {
	config, err := NewConfig([]byte(config))
	if err != nil {
		t.Fatal("Error during config parsing")
	}

	if config.Services[0] != "service1" {
		t.Error("service1 not found")
	}

	if config.Services[1] != "service2" {
		t.Error("service2 not found")
	}
}

func TestNewConfigHipchat(t *testing.T) {
	config, err := NewConfig([]byte(config))
	if err != nil {
		t.Fatal("Error during config parsing")
	}

	hipchat := config.Notifier.Hipchat
	if hipchat == nil {
		t.Fatal("Error during notify config parsing")
	}

	if hipchat.Room != "room" {
		t.Error("Invalid hipchat room")
	}

	if hipchat.Token != "token" {
		t.Error("Invalid hipchat token")
	}

	if !hipchat.Active {
		t.Error("Invalid hipchat on_active")
	}

	if !hipchat.Failed {
		t.Error("Invalid hipchat on_failed")
	}

	if !hipchat.Restart {
		t.Error("Invalid hipchat on_restart")
	}
}

func TestNewConfigSlack(t *testing.T) {
	config, err := NewConfig([]byte(config))
	if err != nil {
		t.Fatal("Error during config parsing")
	}

	slack := config.Notifier.Slack
	if slack == nil {
		t.Fatal("Error during notify config parsing")
	}

	if slack.WebhookURL != "webhook_url" {
		t.Error("Invalid slack webhook_url")
	}

	if slack.Channel != "channel" {
		t.Error("Invalid slack channel")
	}

  if slack.Username != "username" {
    t.Error("Invalid slack username")
  }

	if !slack.Active {
		t.Error("Invalid slack on_active")
	}

	if !slack.Failed {
		t.Error("Invalid slack on_failed")
	}

	if !slack.Restart {
		t.Error("Invalid slack on_restart")
	}
}
