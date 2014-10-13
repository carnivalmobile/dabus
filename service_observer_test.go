package main

import (
	"testing"
	"time"

	"github.com/coreos/go-systemd/dbus"
)

func TestNewServiceObserver(t *testing.T) {
	target := "redis@arch.service"
	err := StartUnit(target)
	if err != nil {
		t.Fatalf("Enable to start unit: %s", err)
	}

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		close(timeout)
	}()

	obs, err := NewServiceObserver([]string{target})
	if err != nil {
		t.Fatalf("Enable to connect to the DUBS: %s", err)
	}

	eventChan := obs.Observe()
	for {
		select {
		case event := <-eventChan:
			expected := ServiceEvent{target, "active", "running"}
			if (*event) != expected {
				t.Fatalf("Unexpected event %#v", event)
			}
			goto success
		case <-timeout:
			t.Fatal("Reached timeout")
		}
	}
success:
	return
}

func StartUnit(target string) error {
	conn, err := dbus.New()
	if err != nil {
		return err
	}

	reschan := make(chan string)
	_, err = conn.StartUnit(target, "replace", reschan)
	if err != nil {
		return err
	}

	job := <-reschan
	if job != "done" {
		return err
	}

	return nil
}
