package models

import (
	messaggio "httpserve/proto"
	"net/http"
)

type Request struct {
	Intents  []*messaggio.MessageSendingIntent `json:"intents"`
	Response *messaggio.BareResponse           `json:"response"` // check to validation
}

func (c *Request) Bind(r *http.Request) error {
	return nil
}

func (p *Request) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
