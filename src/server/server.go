package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/server/middlewares"
	"github.com/gsn_manager_service/src/server/routes"
)

func StartServer(cfg *config.Config) *http.Server {
	r := chi.NewRouter()
	middlewares.SetupMiddleware(r)

	routes.SetupRoutes(r)

	// err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), r)
	// if err != nil {
	// 	adapters.Logger.Error().Msg(fmt.Sprintf("Server failed to start: %v", err))
	// 	return err
	// }
	// return nil

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.PORT),
		Handler: r,
	}

	return server
}
