package cli

import (
	"context"
	version "directory/pkg"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"

	"directory/pkg/database"
	"directory/pkg/logger"
	"directory/pkg/router"
	"directory/pkg/server"
)

type Command struct {
}

func (c *Command) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := load()
	if err != nil {
		logger.Error(ctx, "error loading configuration", err)
		os.Exit(1)
	}

	s := &server.Server{
		Address:           cfg.Server.Address,
		Handler:           router.New(database.New(cfg.Database.Driver, cfg.Database.Source)).Mux,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
	}

	var g errgroup.Group

	g.Go(func() error {
		logger.Info(ctx, fmt.Sprintf("server listening at %s", cfg.Server.Address),
			"revision", version.Latest.CommitHash(),
			"debug", cfg.Debug,
		)
		err := s.ListenAndServe(ctx)

		stop()
		logger.Info(ctx, "server stopped", "err", err)
		return err
	})

	<-ctx.Done()
	logger.Info(ctx, "shutting down server")

	ctx = context.Background()

	if err = g.Wait(); err != nil {
		logger.Error(ctx, "error shutting down server", err)
		os.Exit(1)
	}

	logger.Info(ctx, "server shutdown successfully")

	return err
}
