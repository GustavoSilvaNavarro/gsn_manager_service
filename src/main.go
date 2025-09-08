package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gsn_manager_service/src/adapters"
	"github.com/gsn_manager_service/src/config"
	"github.com/gsn_manager_service/src/connections"
	"github.com/gsn_manager_service/src/server"
)

func main() {
	// Load the application configuration.
	config := config.LoadConfig()
	adapters.InitLogger()

	result, err := connections.StartConnections()
	if err != nil {
		os.Exit(1)
	}
	defer adapters.DisconnectMongo(result.Db)

	connections.CreateAllFactories(result.Db)

	// TODO: for now this root context is not being used, so is not doing anything
	// I could remove it, but it would be used later
	_, cancel := context.WithCancel(context.Background()) // Create root ctx
	defer cancel()

	srv := server.StartServer(config)
	serverErr := make(chan error, 1)
	go func() {
		adapters.Logger.Info().Msg(fmt.Sprintf("🚀 Starting server on %d port", config.PORT))
		if err := srv.ListenAndServe(); err != nil {
			serverErr <- err
		}
	}()

	// Signal listener goroutine
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		adapters.Logger.Info().Msgf("📴 Received shutdown signal: %s", sig)
		cancel()
	case err := <-serverErr:
		adapters.Logger.Error().Err(err).Msg("💥 Server crashed")
		cancel()
	}

	// Give some time for cleanup before force exit
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		adapters.Logger.Error().Err(err).Msg("⚠️ Server forced to shutdown")
	} else {
		adapters.Logger.Info().Msg("✅ Server stopped gracefully")
	}

	adapters.Logger.Info().Msg("👋 Cleanup complete, exiting.")
}
