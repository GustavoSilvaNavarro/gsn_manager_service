package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	NAME                 string
	ENVIRONMENT          string
	PORT                 int
	MONGO_URI            string
	DB_NAME              string
	TASK_COLLECTION_NAME string
}

var Cfg *Config

func LoadConfig() *Config {
	viper.SetConfigFile(".env") // or path to your .env file
	viper.SetConfigType("env")  // dotenv format

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No .env file found, using environment variables / defaults")
	}

	viper.AutomaticEnv()

	viper.SetDefault("NAME", "gsn_expenses_tracker")
	viper.SetDefault("ENVIRONMENT", "dev")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")
	viper.SetDefault("DB_NAME", "table")

	cfg := &Config{
		NAME:                 viper.GetString("NAME"),
		ENVIRONMENT:          viper.GetString("ENVIRONMENT"),
		PORT:                 viper.GetInt("PORT"),
		MONGO_URI:            viper.GetString("MONGO_URI"),
		DB_NAME:              viper.GetString("DB_NAME"),
		TASK_COLLECTION_NAME: "tasks",
	}

	Cfg = cfg
	return cfg
}
