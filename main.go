package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const doc = "Usage: dabus config.yaml"

func main() {
	if len(os.Args) != 2 {
		fmt.Println(doc)
		return
	}

	configFile := os.Args[1]

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Invalid config file: %v\n", err)
		return
	}

	config, err := NewConfig(data)
	if err != nil {
		fmt.Printf("Unable to parse config file: %v\n", err)
		return
	}

	obs, err := NewServiceObserver(config.Services)
	if err != nil {
		fmt.Printf("Unable to connect to the DBUS: %v\n", err)
		return
	}

	eventChan := obs.Observe()
	for {
		select {
		case event := <-eventChan:
			config.Notifier.Send(event)
		}
	}
}
