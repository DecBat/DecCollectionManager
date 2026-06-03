package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DecBat/DecCollectionManager/internal/handlers"
	"github.com/DecBat/DecCollectionManager/internal/routes"
	"github.com/DecBat/DecCollectionManager/serverconfig"
)

func main() {
	// Load config
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
	// Create a new handler
	handler := handlers.NewHandlers()

	// set up the HTTP server
	mux := http.NewServeMux()

	// Set routes
	routes.SetupRoutes(mux, handler)
	// server instance
	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf("Server is up and running on Port%s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
