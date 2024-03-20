package delivery

import (
	"encoding/json"
	messaggio "httpserve/proto"
	"log"
	"net/http"
	"sync"

	"httpserve/models"
	"httpserve/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/go-playground/validator/v10"
	// "github.com/rs/zerolog"
)

var validate = validator.New()

type requestDelivery struct {
	usecase usecase.ClaimMessageSendingRequest
	// logger  zerolog.Logger
}

type Request struct {
	Intents []*messaggio.MessageSendingIntent `json:"intents"`
}

func NewHandler(r chi.Router, usecase usecase.ClaimMessageSendingRequest) {
	handler := requestDelivery{
		usecase: usecase,
		// logger:  logger,
	}
	r.Route("/api/ClaimMessageSending", func(r chi.Router) {
		r.Post("/", handler.handleRequest)
	})
}

func (h *requestDelivery) handleRequest(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{}
	req := &Request{}

	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, req)
		return
	}

	if err := validate.Struct(req); err != nil {
		render.Render(w, r, req)
		return
	}

	if len(req.Intents) > models.LimitRequests {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		msg := &models.Messages{
			ClaimMessages: make(chan *messaggio.MessageSendingIntent, models.LimitRequests),
			BareResponse:  make(chan *models.IntentResponse, models.LimitRequests),
		}

		for _, rq := range req.Intents {
			msg.ClaimMessages <- rq
		}

		h.usecase.CollectionOfRequests(msg.ClaimMessages, msg.BareResponse)

		// Получаем response и закрываем канал
		resp := <-msg.BareResponse
		close(msg.BareResponse)

		respJSON, err := json.Marshal(resp)
		if err != nil {
			log.Printf("error marshalling response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Установка заголовков и отправка JSON-ответа
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJSON)
	}()

	wg.Wait()
}

func (c *Request) Bind(r *http.Request) error {
	return nil
}

func (p *Request) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
