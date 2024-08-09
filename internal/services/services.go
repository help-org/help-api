package services

import "github.com/go-chi/chi/v5"

type Service interface {
	RegisterRoutes(mux *chi.Mux)
}
