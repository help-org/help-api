package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"directory/internal/store/database"
)

type DivisionService struct {
	divisionStore database.DivisionStore
}

func NewDivisionService(divisionsStore database.DivisionStore) *DivisionService {
	return &DivisionService{
		divisionStore: divisionsStore,
	}
}

func (s *DivisionService) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/divisions/{id}", s.FindByID)
}

func (s *DivisionService) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "division id is invalid", http.StatusBadRequest)
		return
	}

	division, err := s.divisionStore.FindByID(ctx, id)
	if err != nil {
		http.Error(w, "division id was not found", http.StatusNotFound)
	}

	response, err := json.Marshal(division)
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
