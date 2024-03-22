package usecase

import (
	messaggio "httpserve/proto"
)

//go:generate mockgen -source=usecase.go -destination/mocks/mock.go

type ClaimMessageSendingRequest interface {
	ClaimMessageSending([]*messaggio.MessageSendingIntent) ([]*messaggio.BareResponse, error)
	SendToChannel([]*messaggio.MessageSendingIntent, chan *messaggio.BareResponse)
	NewClient() messaggio.MessaggioBillingClient
	CollectionOfRequests()
}
