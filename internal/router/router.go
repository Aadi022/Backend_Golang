package router

import (
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/handler"
	"github.com/Aadi022/Backend_Golang/internal/middleware"
)

func New(health *handler.HealthHandler) http.Handler {
	mux := http.NewServeMux()

	healthHandler := middleware.Recovery(
		middleware.Logging(
			http.HandlerFunc(health.Health),
		),
	)

	readyHandler := middleware.Recovery(
		middleware.Logging(
			http.HandlerFunc(health.Ready),
		),
	)

	mux.Handle("/health", healthHandler)
	mux.Handle("/ready", readyHandler)

	return mux
}
