package config

import "github.com/spf13/viper"

type Config struct {
	NAME        string
	ENVIRONMENT string
	PORT        int
	MONGO_URI   string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("NAME", "gsn_expenses_tracker")
	viper.SetDefault("ENVIRONMENT", "dev")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")

	cfg := &Config{
		NAME:        viper.GetString("NAME"),
		ENVIRONMENT: viper.GetString("ENVIRONMENT"),
		PORT:        viper.GetInt("PORT"),
		MONGO_URI:   viper.GetString("MONGO_URI"),
	}

	return cfg
}
