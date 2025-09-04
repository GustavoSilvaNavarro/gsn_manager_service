package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	NAME        string
	ENVIRONMENT string
	PORT        int
}

func LoadConfig() (*Config, error) {
	// Read the NAME environment variable.
	name := os.Getenv("NAME")
	if name == "" {
		name = "gsn_expenses_tracker"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "dev" // Default from the TS example
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port
	}

	serverPort, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid server port: %w", err)
	}

	return &Config{
		NAME:        name,
		ENVIRONMENT: environment,
		PORT:        serverPort,
	}, nil
}
