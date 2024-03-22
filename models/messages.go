package models

import messaggio "httpserve/proto"

var (
	LimitRequests = 10
)

type MessageSendingIntent struct {
	Intents []*messaggio.MessageSendingIntent
	Channel chan *messaggio.BareResponse
}
