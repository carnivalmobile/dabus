package main

type Notifier interface {
	Send(event *ServiceEvent) error
}

type Notification struct {
	Hipchat *Hipchat `yaml:"hipchat,omitempty"`
	Slack   *Slack   `yaml:"slack,omitempty"`
}

func (n *Notification) Send(event *ServiceEvent) error {
	if n.Hipchat != nil {
		event.SendWithNotifier(n.Hipchat)
	}
	if n.Slack != nil {
		event.SendWithNotifier(n.Slack)
	}

	return nil
}
