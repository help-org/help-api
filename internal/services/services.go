package services

import (
	"github.com/go-chi/chi/v5"

	"directory/internal/services/directory"
)

type Service interface {
	RegisterRoutes(mux *chi.Mux)
}

var Services = []Service{
	directory.New(),
}
