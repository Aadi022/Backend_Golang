package router

import (
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/handler"
	"github.com/Aadi022/Backend_Golang/internal/middleware"
	"github.com/Aadi022/Backend_Golang/internal/service"
)

// Constructor function that creates and returns the router
func New(
	health *handler.HealthHandler,
	user *handler.UserHandler,
	jwtService *service.JWTService, // JWT Service required by Auth middleware to validate tokens
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
		"POST /api/v1/auth/login", // Register the login endpoint
		middleware.Recovery(
			middleware.Logging(
				http.HandlerFunc(user.Login),
			),
		),
	)

	// Protected routes
	mux.Handle(
		"GET /api/v1/profile", // Only authenticated users can access this endpoint
		middleware.Recovery(
			middleware.Logging(
				middleware.Auth(jwtService)( // Validate JWT before executing the handler
					http.HandlerFunc(user.Profile),
				),
			),
		),
	)

	return mux
}
