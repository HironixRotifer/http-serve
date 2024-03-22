package usecase

import (
	"context"
	messaggio "httpserve/proto"
	"io"
	"log"
	"time"

	"httpserve/models"
	"httpserve/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client  messaggio.MessaggioBillingClient
	Channel chan *models.MessageSendingIntent
	ClaimMessageSendingRequest
}

func NewUsecase() ClaimMessageSendingRequest {
	ch := make(chan *models.MessageSendingIntent, models.LimitRequests)

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
	message := &models.MessageSendingIntent{
		Intents: intents,
		Channel: channel,
	}

	c.Channel <- message
}

// Складывает Intent-ы и отправляет их каждые 10 секунд или каждые LimitRequests записей
func (c *Client) CollectionOfRequests() {
	flushFn := func(ctx context.Context, requests []*models.MessageSendingIntent) error {
		intents := make([]*messaggio.MessageSendingIntent, 0, models.LimitRequests*models.LimitRequests)

		for _, value := range requests {
			intents = append(intents, value.Intents...)
		}

		resp, err := c.ClaimMessageSending(intents)
		if err != nil {
			log.Printf("Fatal error with sending intents: %v", err)
		}

		// На данный момент алгоритм рассчитан на последовательность элементов в возвращаемом срезе
		count := 0
		for _, v := range requests {
			for i := 0; i < len(v.Intents); i++ {
				v.Channel <- resp[count]
				count += 1
			}
			close(v.Channel)
		}

		return nil
	}

	err := utils.WriteBatch(c.Channel, models.LimitRequests, 10*time.Second, 10*time.Second, flushFn)
	if err != nil {
		log.Fatal("error writing batch ", err)
		return
	}
}

func (c *Client) ClaimMessageSending(Intents []*messaggio.MessageSendingIntent) ([]*messaggio.BareResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	bareResponse := make([]*messaggio.BareResponse, 0)
	defer cancel()

	claim, err := c.Client.ClaimMessageSending(ctx, &messaggio.ClaimMessageSendingRequest{
		Intents: Intents,
	})
	if err != nil {
		return nil, err
	}

	for {
		resp, err := claim.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error while streaming %v", err)
		}
		bareResponse = append(bareResponse, resp)
	}

	return bareResponse, nil
}
