package main

import (
	"log"
	"os"

	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server"
)

func main() {
	// Load the application configuration.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	server.StartServer(config)
}
