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
	Client messaggio.MessaggioBillingClient
	ClaimMessageSendingRequest
}

type ClaimMessageSendingRequest interface {
	NewClient() messaggio.MessaggioBillingClient
	ClaimMessageSending([]*messaggio.MessageSendingIntent) (*messaggio.BareResponse, error)
	CollectionOfRequests(chan *messaggio.MessageSendingIntent, chan *models.IntentResponse)
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

	client := messaggio.NewMessaggioBillingClient(conn)

	return client
}

// Складывает Intent-ы и отправляет их каждые 10 секунд или каждые LimitRequests записей
func (c *Client) CollectionOfRequests(claimMessage chan *messaggio.MessageSendingIntent, bareResponse chan *models.IntentResponse) {

	flushFn := func(ctx context.Context, Intents []*messaggio.MessageSendingIntent) error {
		resp, err := c.ClaimMessageSending(Intents)
		if err != nil {
			return err
		}

		close(claimMessage)
		r := &models.IntentResponse{
			Intents:  Intents,
			Response: resp,
		}

		bareResponse <- r

		return nil
	}

	err := utils.WriteBatch(claimMessage, models.LimitRequests, 10*time.Second, 10*time.Second, flushFn)
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
	stream, err := c.Client.ReportMessageDelivery(context.Background(), &messaggio.ReportMessageDeliveryRequest{
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
