package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	// MyClient api wrapper
	MyClient struct {
		client  *http.Client
		baseURL string
	}
	// Message from View request
	Message struct {
		Text string `json:"message"`
	}
	// ConflictError
	ConflictError struct{}
)

func (ConflictError) Error() string {
	return "Client request is conflict"
}

// NewMyClient pi wrapper
func NewMyClient(client *http.Client, baseURL string) MyClient {
	return MyClient{
		client:  client,
		baseURL: baseURL,
	}
}

// ViewMessage from the server
func (m MyClient) ViewMessage() (*Message, error) {
	var msg Message
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/view", m.baseURL), nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, ConflictError{}
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
