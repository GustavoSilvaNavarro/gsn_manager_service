package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gsn_manager_service/src/adapters"
	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server/routes"
)

func StartServer(cfg *config.Config) {
	// new server
	// mux := http.NewServeMux()
	r := chi.NewRouter()

	// Routes
	routes.SetupRoutes(r)

	adapters.Logger.Info().Msg(fmt.Sprintf("🚀 Starting server on %d port", cfg.PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), r)
	if err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Server failed to start: %v", err))
		os.Exit(1)
	}
}
