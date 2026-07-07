package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Aadi022/Backend_Golang/internal/config"
	"github.com/Aadi022/Backend_Golang/internal/database"
	"github.com/Aadi022/Backend_Golang/internal/handler"
	"github.com/Aadi022/Backend_Golang/internal/repository"
	"github.com/Aadi022/Backend_Golang/internal/router"
	"github.com/Aadi022/Backend_Golang/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to PostgreSQL
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run database migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// ---------------- Dependency Injection ----------------

	userRepo := repository.NewUserRepository(db)                // Create the User Repository
	jwtService := service.NewJWTService(cfg.JWTSecret)          // Create the JWT service using the secret from config
	userService := service.NewUserService(userRepo, jwtService) // Inject Repository and JWT Service into UserService
	userHandler := handler.NewUserHandler(userService)          // Inject UserService into UserHandler

	healthHandler := handler.NewHealthHandler() // Create Health Handler

	// ------------------------------------------------------

	r := router.New(
		healthHandler,
		userHandler,
	)

	server := &http.Server{ // This object represents your HTTP server and lets you configure and control it.
		Addr:         ":" + cfg.Port,
		Handler:      r,                 // Whenever a request arrives, pass it to this router.
		ReadTimeout:  5 * time.Second,   // The client has 5 seconds to send its request.
		WriteTimeout: 10 * time.Second,  // The server has 10 seconds to finish writing the response.
		IdleTimeout:  120 * time.Second, // Close idle connections after 120 seconds.
	}

	log.Printf("Starting server on :%s", cfg.Port)

	go func() { // Start the HTTP server in a separate goroutine

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Ignore ErrServerClosed because it occurs during graceful shutdown
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)   // Create a channel to receive OS interrupt signals
	signal.Notify(quit, os.Interrupt) // Notify this channel when Ctrl+C (SIGINT) is pressed

	<-quit // Block until an interrupt signal is received

	log.Println("Shutting down the server...")

	// A context tells a function how long it's allowed to run and when it should stop.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Create a context with a 10-second timeout

	/*
		ctx    -> Carries the 10-second timeout and is passed to server.Shutdown().
		cancel -> Cancels the context manually and frees its internal resources.
	*/

	defer cancel() // Ensure resources associated with the context are released

	if err := server.Shutdown(ctx); err != nil { // Gracefully stop accepting new requests and wait for ongoing ones to finish
		log.Fatal(err)
	}

	log.Println("Server stopped successfully")
}
