package models

import messaggio "httpserve/proto"

var (
	LimitRequests = 10
)

type Messages struct {
	ClaimMessages chan *messaggio.MessageSendingIntent
	BareResponse  chan *IntentResponse
}
