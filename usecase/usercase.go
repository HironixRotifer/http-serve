package usecase

import (
	"context"
	"fmt"
	messaggio "httpserve/proto"
	"log"
	"time"

	"httpserve/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Client struct {
	Client messaggio.MessaggioBillingClient
	ClaimMessageSendingRequest
}

type ClaimMessageSendingRequest interface {
	NewClient() messaggio.MessaggioBillingClient
	ClaimMessageSending(models.Intents) (*messaggio.BareResponse, error)
}

func NewUsecase() ClaimMessageSendingRequest {
	return &Client{
		Client: NewClient(),
	}
}

func NewClient() messaggio.MessaggioBillingClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	// defer conn.Close()
	fmt.Println(conn)

	client := messaggio.NewMessaggioBillingClient(conn)

	return client
}

func (c *Client) ClaimMessageSending(Intents models.Intents) (*messaggio.BareResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	claim, err := c.Client.ClaimMessageSending(ctx, &messaggio.ClaimMessageSendingRequest{
		Intents: []*messaggio.MessageSendingIntent{
			{
				Service:        Intents.Service,
				ServiceItem:    Intents.ServiceItem,
				Country:        Intents.Country,
				Operator:       Intents.Operator,
				ProjectId:      Intents.ProjectID,
				GateId:         Intents.GateID,
				GateSenderId:   Intents.GateSenderID,
				ClientSenderId: Intents.ClientSenderID,
				MessageId:      Intents.MessageID,
				PhoneNumber:    Intents.PhoneNumber,
				DeliveryTtl:    Intents.DeliveryTTL,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	defer claim.CloseSend()

	resp, err := claim.Recv()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Так понимаю, что эта реализация должна лежать на сервере биллинга
func (c *Client) ReportMessageDelivery(Intents models.Intents) (*messaggio.BareResponse, error) {
	stream, err := c.Client.ReportMessageDelivery(context.Background(), &messaggio.ReportMessageDeliveryRequest{
		Reports: []*messaggio.MessageDeliveryReport{
			{
				Service:        Intents.Service,
				ServiceItem:    Intents.ServiceItem,
				Country:        Intents.Country,
				Operator:       Intents.Operator,
				ProjectId:      Intents.ProjectID,
				GateId:         Intents.GateID,
				GateSenderId:   Intents.GateSenderID,
				ClientSenderId: Intents.ClientSenderID,
				MessageId:      Intents.MessageID,
				PhoneNumber:    Intents.PhoneNumber,
				SessionId:      "",
				SessionCreated: &wrapperspb.BoolValue{},
				DeliveryStatus: 1,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	resp, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
