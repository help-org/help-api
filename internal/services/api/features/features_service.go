package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"directory/internal/store/database"
)

type FeatureService struct {
	featureStore database.FeatureStore
}

func NewFeatureService(featuresStore database.FeatureStore) *FeatureService {
	return &FeatureService{
		featureStore: featuresStore,
	}
}

func (s *FeatureService) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/features/{id}", s.FindByID)
}

// TODO add query param for find 1 or recursive
func (s *FeatureService) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "feature id is invalid", http.StatusBadRequest)
		return
	}

	features, err := s.featureStore.FindRelationsByID(ctx, id)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}

	response, err := json.Marshal(features)
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
