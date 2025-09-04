package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server/routes"
)

func StartServer(cfg *config.Config) {
	// This is part of the standard library and avoids external dependencies.
	mux := http.NewServeMux()

	// Call the function from the routes package to set up all the endpoints.
	routes.SetupRoutes(mux)

	// Log the server start message.
	log.Printf("🚀 Starting server on %d port", cfg.PORT)

	// Start the server and listen for incoming requests.
	// The server uses the ServeMux to handle requests based on the registered routes.
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
