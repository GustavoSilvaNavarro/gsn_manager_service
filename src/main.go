package main

import (
	"os"

	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/connections"
	"github.com/gsn_manager_service/src/server"
)

func main() {
	// Load the application configuration.
	config := config.LoadConfig()
	result, err := connections.StartConnections()
	if err != nil {
		os.Exit(1)
	}

	connections.CreateAllFactories(result.Db)

	server.StartServer(config)
}
