package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/HectorMRC/backend-tech-test/metrics"
	"github.com/HectorMRC/backend-tech-test/ride"
)

const httpPort = 8080

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/rides", ride.RideStartHandler)
	r.Post("/rides/{id}/finish", ride.RideFinishHandler)
	r.Post("/metrics", metrics.MetricsHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Error starting http server <%s>", err)
	}
}
