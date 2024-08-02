package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"directory/internal/services/directory"
)

type routes map[string]map[string]http.HandlerFunc
type methods map[string]http.HandlerFunc

type Router struct {
	Mux *chi.Mux
}

func New(db *sql.DB) (router *Router) {
	router = &Router{
		Mux: chi.NewRouter(),
	}

	router.Mux.Use(middleware.Logger)

	directoryService := directory.NewService(db)

	router.Mux.Get("/countries/{country}/states/{state}/cities/{city}", directoryService.ListLocal)

	return
}
