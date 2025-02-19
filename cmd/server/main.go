package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"statsbuddy/api"
	"statsbuddy/config"
	"statsbuddy/service"
)

func main() {
	// Configuration flags
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoDBURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	// Load configuration
	config := config.GetConfig()

	// Create service with config
	chatbotService, err := service.NewChatbotService(config, mongoClient)
	if err != nil {
		log.Fatalf("Failed to create chatbot service: %v", err)
	}

	// Create HTTP handler
	handler := api.NewChatbotHandler(chatbotService)

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/api/query", handler.HandleQuery)

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", *port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Graceful shutdown
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}
}
