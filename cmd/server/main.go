package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/HectorMRC/backend-tech-test/metrics"
	"github.com/HectorMRC/backend-tech-test/ride"
)

const (
	HTTP_PORT        = 8080
	ENV_UNLOCK_PRICE = "UNLOCK_PRICE"
	ENV_MINUTE_PRICE = "MINUTE_PRICE"
)

var (
	UnlockPrice = 100
	MinutePrice = 18
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Warn("no dotenv file has been found",
			zap.Error(err))
	}

	if price, exists := os.LookupEnv(ENV_UNLOCK_PRICE); exists {
		price, err := strconv.Atoi(price)
		if err != nil {
			logger.Fatal("parsing string to int",
				zap.Error(err))
		}

		UnlockPrice = price
	}

	if price, exists := os.LookupEnv(ENV_MINUTE_PRICE); exists {
		price, err := strconv.Atoi(price)
		if err != nil {
			logger.Fatal("parsing string to int",
				zap.Error(err))
		}

		MinutePrice = price
	}

	metricsRepo := metrics.NewInMemoryMetricsRepository()
	metricsApp := metrics.NewMetricsApplication(metricsRepo, logger)
	metricsServer := metrics.NewMetricsServer(metricsApp, logger)

	rideRepo := ride.NewInMemoryRideRepository()
	rideApp := ride.NewRideApplication(UnlockPrice, MinutePrice, rideRepo, logger)
	rideServer := ride.NewRideServer(rideApp, logger)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/rides", rideServer.RideStartHandler)
	r.Post("/rides/{id}/finish", rideServer.RideFinishHandler)
	r.Get("/metrics", metricsServer.MetricsHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", HTTP_PORT), r); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Error starting http server <%s>", err)
	}
}
