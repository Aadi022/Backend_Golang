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
)

func main() {
	//Load configuration
	cfg := config.Load()

	db, err := database.Connect(cfg) //connect to PostgreSQL
	if err != nil {
		log.Fatal(err)
	}

	// Run database migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	//Dependency Injection
	userRepo := repository.NewUserRepository(db)
	_ = userRepo //We know that this variable is not used anywhere

	healthHandler := handler.NewHealthHandler()

	r := router.New(healthHandler)

	server := &http.Server{ //This object represents your HTTP server and lets you configure and control it.
		Addr:         ":" + cfg.Port,
		Handler:      r,                //Whenever a request arrives, pass it to this router.
		ReadTimeout:  5 * time.Second,  //The client has 5 seconds to send its request.
		WriteTimeout: 10 * time.Second, //The server has 10 seconds to finish writing the response.
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting server on :%s", cfg.Port)

	go func() { //function to start the server

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { //we dont stop for http.ErrServerClosed as we are entertaining graceful shutdown
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)   //create a channel to capture os interrupt
	signal.Notify(quit, os.Interrupt) //enqueue in the channel for any os interrupt

	<-quit //blocking statement that waits till something arrives in the channel

	log.Println("Shutting down the server...")
	//A context tells a function how long it's allowed to run and when it should stop.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //creates a context with 10 second timeout
	//context.Background() creates an empty context, context.withTimeout creates a child context that is automatically cancelled after specific duration
	/*
		ctx → A context object that carries the 10-second timeout and is passed to server.Shutdown() so it knows how long it has to shut down.
		cancel → A function (func()) that manually cancels the context and frees its internal resources before the timeout if you're done with it.
	*/
	defer cancel()

	if err := server.Shutdown(ctx); err != nil { //starts a graceful shutdown
		log.Fatal(err)
	}

	log.Println("Server stopped successfully")

}
