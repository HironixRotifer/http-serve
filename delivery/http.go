package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"httpserve/models"
	"httpserve/usecase"

	messaggio "httpserve/proto"

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

	req := &models.Request{}
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

	ch := make(chan *messaggio.BareResponse, models.LimitRequests)

	h.usecase.SendToChannel(req.Intents, ch)

	// Получаем response и закрываем канал
	resp := <-ch
	close(ch)

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error marshalling response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}
