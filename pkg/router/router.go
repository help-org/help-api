package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"directory/internal/services/directory"
)

type Router struct {
	Mux *chi.Mux
}

func New(db *sql.DB) (router *Router) {
	router = &Router{
		Mux: chi.NewRouter(),
	}

	router.Mux.Use(middleware.Logger)

	directoryService := directory.NewService(db)
	directoryService.RegisterRoutes(router.Mux)

	return
}
