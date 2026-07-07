package router

import (
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/handler"
	"github.com/Aadi022/Backend_Golang/internal/middleware"
)

// Constructor function that creates and returns the router
func New(
	health *handler.HealthHandler,
	user *handler.UserHandler,
) http.Handler {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Health routes
	mux.Handle(
		"GET /health",
		middleware.Recovery(
			middleware.Logging(
				http.HandlerFunc(health.Health),
			),
		),
	)

	mux.Handle(
		"GET /ready",
		middleware.Recovery(
			middleware.Logging(
				http.HandlerFunc(health.Ready),
			),
		),
	)

	// Authentication routes
	mux.Handle(
		"POST /api/v1/auth/register",
		middleware.Recovery(
			middleware.Logging(
				http.HandlerFunc(user.Register),
			),
		),
	)

	mux.Handle(
		"POST /api/v1/auth/login", //Register the login endpoint
		middleware.Recovery(
			middleware.Logging(
				http.HandlerFunc(user.Login),
			),
		),
	)

	return mux
}
