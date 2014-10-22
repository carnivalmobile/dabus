package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPClient interface {
	PostJSON(url string, data interface{}) error
}

type NotifierHTTPClient struct {
	http.Client
}

func NewNotifierHTTPClient() *NotifierHTTPClient {
	return &NotifierHTTPClient{http.Client{}}
}

func (client *NotifierHTTPClient) PostJSON(url string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
