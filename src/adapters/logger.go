package adapters

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gsn_manager_service/src/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	loggerOnce sync.Once
	Logger     zerolog.Logger
)

func InitLogger() {
	loggerOnce.Do(func() {
		// Configure time format
		zerolog.TimeFieldFormat = time.RFC3339

		// Set log level based on environment
		logLevel := strings.ToLower(config.Cfg.LOG_LEVEL)
		switch logLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		// Configure output format based on environment
		env := strings.ToLower(config.Cfg.ENVIRONMENT)
		if env != "dev" && env != "stg" && env != "prd" {
			// Pretty console output for development
			Logger = log.Output(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
				NoColor:    false,
			})
		} else {
			// Structured JSON output for production
			Logger = zerolog.New(os.Stderr).With().
				Timestamp().
				Str("service", "gsn-manager-service").
				Logger()
		}

		Logger.Info().
			Str("level", zerolog.GlobalLevel().String()).
			Str("environment", env).
			Msg("Logger initialized")
	})
}
