package server

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"directory/pkg/logger"
)

type Server struct {
	Address           string
	Handler           http.Handler
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	server := &http.Server{
		Addr:              s.Address,
		Handler:           s.Handler,
		TLSConfig:         s.TLSConfig,
		ReadTimeout:       s.ReadTimeout,
		ReadHeaderTimeout: s.ReadHeaderTimeout,
		WriteTimeout:      s.WriteTimeout,
		IdleTimeout:       s.IdleTimeout,
		MaxHeaderBytes:    s.MaxHeaderBytes,
	}
	go func() {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logger.Error(ctx, "graceful HTTP server shutdown failed", err)
			}
		}
	}()

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
