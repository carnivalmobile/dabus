package main

import (
	"fmt"

	"github.com/coreos/go-systemd/dbus"
)

type ServiceObserver struct {
	subsSet *dbus.SubscriptionSet
}

type ServiceEvent struct {
	Service      string
	ActiveStatus string
	SubStatus    string
}

func NewServiceObserver(services []string) (*ServiceObserver, error) {
	conn, err := dbus.New()
	if err != nil {
		return nil, err
	}

	subsSet := conn.NewSubscriptionSet()
	for _, service := range services {
		subsSet.Add(service)
	}

	return &ServiceObserver{subsSet}, nil
}

func (obs *ServiceObserver) Observe() chan *ServiceEvent {
	serviceEventChan := make(chan *ServiceEvent)
	evChan, errChan := obs.subsSet.Subscribe()
	go func() {
		for {
			select {
			case changes := <-evChan:
				for service, event := range changes {
					serviceEvent := &ServiceEvent{
						service,
						event.ActiveState,
						event.SubState,
					}
					serviceEventChan <- serviceEvent
				}
			case err := <-errChan:
				fmt.Printf("DBUS connection error %v", err)
			}
		}
	}()

	return serviceEventChan
}
