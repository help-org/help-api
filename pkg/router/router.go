package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() (router *chi.Mux) {
	router = chi.NewRouter()

	router.Use(middleware.Logger)

	return
}
