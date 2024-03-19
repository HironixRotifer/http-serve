package usecase

import (
	"context"
	"fmt"
	messaggio "httpserve/proto"
	"log"
	"time"

	"httpserve/models"
	"httpserve/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Client struct {
	Client messaggio.MessaggioBillingClient
	ClaimMessageSendingRequest
	Messages Messages
}

type Messages struct {
	ClaimMessages chan *messaggio.MessageSendingIntent
	BareResponse  chan *messaggio.BareResponse
}

type ClaimMessageSendingRequest interface {
	NewClient() messaggio.MessaggioBillingClient
	ClaimMessageSending([]*messaggio.MessageSendingIntent) (*messaggio.BareResponse, error)
	SendMessage(Intent *messaggio.MessageSendingIntent)
	CollectionOfRequests()
}

func NewUsecase() ClaimMessageSendingRequest {
	cm := make(chan *messaggio.MessageSendingIntent, 10)
	br := make(chan *messaggio.BareResponse, 10)

	return &Client{
		Client: NewClient(),
		Messages: Messages{
			cm,
			br,
		},
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

// Кладём Intent в канал
func (c *Client) SendMessage(Intent *messaggio.MessageSendingIntent) {
	// Стоит добавлять проверку? CollectionOfRequests отправляет пачку из 10 записей или по таймауту в 10 секунд
	c.Messages.ClaimMessages <- Intent
}

func (c *Client) GetResponse() *messaggio.BareResponse {
	return <-c.Messages.BareResponse
}

// Складывает Intent-ы и отправляет их каждые 10 секунд или каждые 10 записей
func (c *Client) CollectionOfRequests() {

	flushFn := func(ctx context.Context, Intents []*messaggio.MessageSendingIntent) error {
		resp, err := c.ClaimMessageSending(Intents)
		if err != nil {
			return err
		}
		c.Messages.BareResponse <- resp

		return nil
	}

	err := utils.WriteBatch(c.Messages.ClaimMessages, 10, 10*time.Second, 10*time.Second, flushFn)
	if err != nil {
		log.Fatal("error writing batch ", err)
		return
	}
}

func (c *Client) ClaimMessageSending(Intents []*messaggio.MessageSendingIntent) (*messaggio.BareResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	claim, err := c.Client.ClaimMessageSending(ctx, &messaggio.ClaimMessageSendingRequest{
		Intents: Intents,
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
