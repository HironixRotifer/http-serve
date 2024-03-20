package models

import (
	messaggio "httpserve/proto"
)

type IntentResponse struct {
	Intents  []*messaggio.MessageSendingIntent `json:"intent"`
	Response *messaggio.BareResponse           `json:"response"`
}
