package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server/routes"
)

func StartServer(cfg *config.Config) {
	// new server
	mux := http.NewServeMux()

	// Routes
	routes.SetupRoutes(mux)

	log.Printf("🚀 Starting server on %d port", cfg.PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
		os.Exit(1)
	}
}
