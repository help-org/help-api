package directory

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	database *sql.DB
}

func New() *Service {
	return &Service{}
}

func (s *Service) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/countries/{country}/states/{state}/cities/{city}", s.ListLocal)
}

func (s *Service) ListLocal(w http.ResponseWriter, r *http.Request) {
	directory := &Directory{
		Country: chi.URLParam(r, "country"),
		State:   chi.URLParam(r, "state"),
		City:    chi.URLParam(r, "city"),
		Listings: []*Listing{
			{
				Type:  POLICE,
				Name:  "Local Police",
				Phone: "911",
			},
		},
		Ads: []*Ad{
			{
				Type:  LAWYER,
				Name:  "Local Lawyer",
				Phone: "555-555-5555",
			},
		},
	}

	response, err := json.Marshal(directory)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
