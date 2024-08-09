package cli

import (
	"context"
	"directory/internal/router"
	"directory/internal/services"
	"directory/internal/services/api"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"directory/internal/store/database"
	"directory/pkg/config"
	db "directory/pkg/database"
	"directory/pkg/logger"
	"directory/pkg/server"
	"directory/pkg/version"
)

type Command struct {
}

func (c *Command) Serve() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := &config.Config{}
	err := cfg.FromEnv()
	if err != nil {
		logger.Error(ctx, "error loading configuration", err)
		os.Exit(1)
	}

	s, err := initialize(cfg)

	var g errgroup.Group
	g.Go(func() error {
		defer stop()
		logger.Info(ctx, fmt.Sprintf("server listening at %s", cfg.Server.Address),
			"revision", version.Latest.CommitHash(),
			"debug", cfg.Debug,
		)
		err := s.ListenAndServe(ctx)

		logger.Info(ctx, "server stopped", "err", err)
		return err
	})

	<-ctx.Done()
	ctx = context.Background()

	if err = g.Wait(); err != nil {
		logger.Error(ctx, "error shutting down server", err)
		os.Exit(1)
	}

	logger.Info(ctx, "server shutdown successfully")

	return err
}

func initialize(cfg *config.Config) (s *server.Server, err error) {
	db, err := db.Connect(cfg)

	var services []services.Service

	divisionStore := database.NewDivisionStore(db)
	divisionService := api.NewDivisionService(*divisionStore)
	services = append(services, divisionService)

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
