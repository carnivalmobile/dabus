package main

type Notifier interface {
	Send(event *ServiceEvent) error
}

type Notification struct {
	Hipchat *Hipchat `yaml:"hipchat,omitempty"`
}

func (n *Notification) Send(event *ServiceEvent) error {
	event.SendWithNotifier(n.Hipchat)

	return nil
}
