package main

import (
	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server"
)

func main() {
	// Load the application configuration.
	config := config.LoadConfig()

	server.StartServer(config)
}
