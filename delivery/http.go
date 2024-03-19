package delivery

import (
	"fmt"
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

type requestDelivery struct {
	usecase usecase.ClaimMessageSendingRequest
	// logger  zerolog.Logger
}

type Request struct {
	Intents []models.Intents `json:"intents"`
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
	var validate = validator.New()
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

	fmt.Println(req) // TODO: удалить

	for _, intents := range req.Intents {
		wg.Add(1)
		go func(intent models.Intents) {
			defer wg.Done()
			resp, err := h.usecase.ClaimMessageSending(intent)
			if err != nil {
				log.Printf("error sending request: %v", err)
			}
			log.Printf("Get response: %v", resp)
		}(intents)
	}

	render.Status(r, http.StatusOK)
}

func (c *Request) Bind(r *http.Request) error {
	return nil
}

func (p *Request) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
