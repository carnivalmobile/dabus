package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Services []string      `yaml:"services,omitempty"`
	Notifier *Notification `yaml:"notify,omitempty"`
}

func NewConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config, nil
}
