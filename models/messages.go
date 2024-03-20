package models

import messaggio "httpserve/proto"

var (
	LimitRequests = 10
)

type Messages struct {
	Intents  []*messaggio.MessageSendingIntent
	Channels chan *messaggio.BareResponse
}
