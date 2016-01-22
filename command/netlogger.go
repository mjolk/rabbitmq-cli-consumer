package command

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ProvisionEvent struct {
	Error   bool   `json:"is_error"`
	Output  string `json:"out"`
	Message []byte `json:"message"`
}

type NetLogger struct {
	Address string
}

func (n *NetLogger) SendOK(p []byte, bod []byte) error {
	event := ProvisionEvent{
		Error:   false,
		Output:  string(p),
		Message: bod,
	}
	return n.send(&event)
}

func (n *NetLogger) SendError(p []byte, bod []byte) error {
	event := ProvisionEvent{
		Error:   true,
		Output:  string(p),
		Message: bod,
	}
	return n.send(&event)
}

func (n *NetLogger) send(p *ProvisionEvent) error {
	post, _ := json.Marshal(p)
	_, err := http.Post(n.Address, "encoding/json", bytes.NewBuffer(post))
	if err != nil {
		return err
	}
	return nil

}
