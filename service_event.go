package main

import "log"

type ServiceEvent struct {
	Service      string
	ActiveStatus string
	SubStatus    string
}

func (event *ServiceEvent) SendWithNotifier(notifier Notifier) error {
	if notifier == nil {
		return nil
	}

	err := notifier.Send(event)
	if err != nil {
		log.Println(err)
	}

	return err
}
