package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	backend "github.com/HectorMRC/backend-tech-test"
)

const httpPort = 8080

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/rides", backend.RideStartHandler)
	r.Post("/rides/{id}/finish", backend.RideFinishHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Error starting http server <%s>", err)
	}
}
