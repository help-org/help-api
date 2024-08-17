package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"directory/pkg/config"
	"directory/pkg/logger"
)

type connection struct {
	DB Pool

	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Pool interface {
	Ping(ctx context.Context) (err error)
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, arguments ...any) (rows pgx.Rows, err error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
}

func (c *connection) Querier() Pool {
	return c.DB
}

func (c *connection) Ping(ctx context.Context) (err error) {
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = c.DB.Exec(pingCtx, ";")
	if err != nil {
		logger.Warn(ctx, "database ping failed", err)
	}
	return
}

func (c *connection) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return c.DB.Exec(ctx, sql, arguments...)
}

func (c *connection) Query(ctx context.Context, sql string, arguments ...any) (rows pgx.Rows, err error) {
	return c.DB.Query(ctx, sql, arguments...)
}

func (c *connection) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	return c.DB.QueryRow(ctx, sql, arguments...)
}

func Connect(cfg *config.Config) (pool Pool, err error) {
	ctx := context.Background()

	pgxConfig, err := pgxpool.ParseConfig(cfg.Database.Source)
	if err != nil {
		return
	}
	pgxConfig.MaxConnLifetime = cfg.Database.MaxConnLifetime
	pgxConfig.MaxConns = int32(cfg.Database.MaxConnections)
	pgxConfig.ConnConfig.ConnectTimeout = cfg.Database.ConnectTimeout

	connect := &connection{
		readTimeout:  cfg.Database.ReadTimeout,
		writeTimeout: cfg.Database.WriteTimeout,
	}
	connect.DB, err = pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for {
		err = connect.Ping(ctx)
		if err == nil || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			break
		}
		time.Sleep(time.Second)
	}

	pool = connect
	logger.Info(ctx, "database connected")

	return
}
