package database

import (
	"context"
	"directory/pkg/config"
	"directory/pkg/logger"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
	"time"
)

var instance Pool
var mu sync.Mutex

type db struct {
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

func (p *db) Querier() Pool {
	return p.DB
}

func (p *db) Ping(ctx context.Context) (err error) {
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = p.DB.Exec(pingCtx, ";")
	if err != nil {
		logger.Warn(ctx, "database ping failed", err)
	}
	return
}

func (p *db) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.DB.Exec(ctx, sql, arguments...)
}

func (p *db) Query(ctx context.Context, sql string, arguments ...any) (rows pgx.Rows, err error) {
	return p.DB.Query(ctx, sql, arguments...)
}

func (p *db) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	return p.DB.QueryRow(ctx, sql, arguments...)
}

func Connect(cfg *config.Config) (pool Pool, err error) {
	ctx := context.Background()

	mu.Lock()
	defer mu.Unlock()

	if instance != nil {
		pool = instance
		return
	}

	s := &db{
		readTimeout:  cfg.Database.ReadTimeout,
		writeTimeout: cfg.Database.WriteTimeout,
	}
	pool = s

	pgxConfig, err := pgxpool.ParseConfig(cfg.Database.Source)
	if err != nil {
		return
	}

	pgxConfig.MaxConnLifetime = cfg.Database.MaxConnLifetime
	pgxConfig.MaxConns = int32(cfg.Database.MaxConnections)
	pgxConfig.ConnConfig.ConnectTimeout = cfg.Database.ConnectTimeout

	s.DB, err = pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return
	}

	err = ping(s)
	if err != nil {
		return
	}
	instance = pool

	logger.Info(ctx, "database connected")
	return
}

func ping(s Pool) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		err = s.Ping(ctx)
		if err == nil || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return
		}
		time.Sleep(time.Second)
	}
}
