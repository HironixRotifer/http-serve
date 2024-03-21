package usecase

import (
	"context"
	messaggio "httpserve/proto"
	"log"
	"time"

	"httpserve/models"
	"httpserve/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client  messaggio.MessaggioBillingClient
	Channel chan *models.Messages
	ClaimMessageSendingRequest
}

type ClaimMessageSendingRequest interface {
	ClaimMessageSending([]*messaggio.MessageSendingIntent) (*messaggio.BareResponse, error)
	SendToChannel([]*messaggio.MessageSendingIntent, chan *messaggio.BareResponse)
	NewClient() messaggio.MessaggioBillingClient
	CollectionOfRequests()
}

func NewUsecase() ClaimMessageSendingRequest {
	ch := make(chan *models.Messages, models.LimitRequests)

	return &Client{
		Client:  NewClient(),
		Channel: ch,
	}
}

func NewClient() messaggio.MessaggioBillingClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	// defer conn.Close()

	client := messaggio.NewMessaggioBillingClient(conn)

	return client
}

func (c *Client) SendToChannel(intents []*messaggio.MessageSendingIntent, channel chan *messaggio.BareResponse) {
	messageIntent := make([]*messaggio.MessageSendingIntent, models.LimitRequests)
	message := &models.Messages{
		Intents:  messageIntent,
		Channels: channel,
	}

	c.Channel <- message
}

// Складывает Intent-ы и отправляет их каждые 10 секунд или каждые LimitRequests записей
func (c *Client) CollectionOfRequests() {
	flushFn := func(ctx context.Context, requests []*models.Messages) error {
		// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		// defer cancel()

		intents := make([]*messaggio.MessageSendingIntent, 0, models.LimitRequests*models.LimitRequests)
		channels := make([]chan *messaggio.BareResponse, 0, len(requests))

		for _, value := range requests {
			intents = append(intents, value.Intents...)
			channels = append(channels, value.Channels)
		}

		resp, err := c.ClaimMessageSending(intents)
		if err != nil {
			resp = &messaggio.BareResponse{
				Error: &messaggio.Error{
					Code:    "500",
					Message: "Internal Server Error",
				},
			}
		}

		for _, ch := range channels {
			select {

			case ch <- resp:

			case <-time.After(time.Millisecond):

			}
		}

		return nil
	}

	err := utils.WriteBatch(c.Channel, models.LimitRequests, 10*time.Second, 10*time.Second, flushFn)
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

func (c *Client) ReportMessageDelivery(Intents []*messaggio.MessageDeliveryReport) (*messaggio.BareResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := c.Client.ReportMessageDelivery(ctx, &messaggio.ReportMessageDeliveryRequest{
		Reports: Intents,
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
