package router

import (
	"net/http"

	"github.com/Aadi022/Backend_Golang/internal/handler"
)

// Constructor function that creates and returns the router
func New(health *handler.HealthHandler) http.Handler {
	//http.NewServeMux() creates a new HTTP request multiplexer (router) that maps URL paths to their corresponding handler functions.
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Health)
	mux.HandleFunc("/ready", health.Ready)
	return mux
}
