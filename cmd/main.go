package main

import (
	"fmt"
	"httpserve/delivery"
	messaggio "httpserve/proto"
	"httpserve/usecase"

	"net/http"

	"github.com/go-chi/chi"
	// "google.golang.org/grpc"
)

var (
	port       = 8080
	IntentChan = make(chan *messaggio.MessageSendingIntent, 10)
)

func main() {
	var r = chi.NewRouter()

	uc := usecase.NewUsecase()
	delivery.NewHandler(r, uc)

	// Запуск горутины, которая будет собирать сообщения
	go uc.CollectionOfRequests()

	fmt.Printf("HTTP server listening on %v \n", port)
	server := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: r}
	server.ListenAndServe()
}
