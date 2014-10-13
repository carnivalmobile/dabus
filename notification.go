package main

import "log"

type Notification struct {
	Hipchat *Hipchat `yaml:"hipchat,omitempty"`
}

func (n *Notification) Send(event *ServiceEvent) error {
	if n.Hipchat != nil {
		err := n.Hipchat.Send(event)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
