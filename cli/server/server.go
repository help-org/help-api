package server

import (
	"context"
	features "directory/internal/services/api/features"
	listings "directory/internal/services/api/listings"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"directory/internal/router"
	"directory/internal/services"
	"directory/internal/store/database"
	"directory/pkg/config"
	db "directory/pkg/database"
	"directory/pkg/logger"
	"directory/pkg/server"
)

func Serve() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := &config.Config{}
	err := cfg.FromEnv()
	if err != nil {
		logger.Error(ctx, "error loading configuration", err)
		os.Exit(1)
	}

	s, err := initialize(cfg)
	if err != nil {
		logger.Error(ctx, "error initializing server", err)
		os.Exit(1)
	}

	logger.Info(ctx, fmt.Sprintf("server listening at %s", cfg.Server.Address))
	err = s.ListenAndServe(ctx)
	if err != nil {
		logger.Error(ctx, "error starting server", err)
		os.Exit(1)
	}

	logger.Info(ctx, "server shutdown successfully")

	return err
}

func initialize(cfg *config.Config) (s *server.Server, err error) {
	db, err := db.Connect(cfg)

	var services []services.Service
	// Stores
	featureStore := database.NewFeatureStore(db)
	listingStore := database.NewListingStore(db)

	// Services
	featuresService := features.NewFeatureService(*featureStore)
	listingService := listings.NewListingService(*featureStore, *listingStore)

	services = append(services, featuresService, listingService)

	var middlewares chi.Middlewares
	middlewares = append(middlewares, middleware.Logger)
	mux := router.NewMuxer(middlewares, services)
	handler := router.New(mux)

	s = &server.Server{
		Address:           cfg.Server.Address,
		Handler:           handler,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
	}

	return
}
