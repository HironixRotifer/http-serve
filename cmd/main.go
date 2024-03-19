package main

import (
	"fmt"
	"httpserve/delivery"
	"httpserve/usecase"
	"net/http"

	"github.com/go-chi/chi"
	// "google.golang.org/grpc"
)

var (
	port = 8080
)

func main() {
	var r = chi.NewRouter()

	uc := usecase.NewUsecase()
	delivery.NewHandler(r, uc)

	fmt.Printf("HTTP server listening on %v \n", port)
	server := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: r}
	server.ListenAndServe()
}
