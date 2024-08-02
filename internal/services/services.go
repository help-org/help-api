package services

import "github.com/go-chi/chi/v5"

type Service interface {
	RegisterRoutes(r chi.Router)
}
