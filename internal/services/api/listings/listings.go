package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"directory/internal/store/database"
)

type ListingService struct {
	featureStore database.FeatureStore
	listingStore database.ListingStore
}

func NewListingService(featureStore database.FeatureStore, listingStore database.ListingStore) *ListingService {
	return &ListingService{
		featureStore: featureStore,
		listingStore: listingStore,
	}
}

func (s *ListingService) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/frontend/feature/{featureId}/listing", s.FindListingsByFeatureId)
}

// TODO add query param for find 1 or recursive
func (s *ListingService) FindListingsByFeatureId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	featureId, err := strconv.Atoi(chi.URLParam(r, "featureId"))
	if err != nil {
		http.Error(w, "feature id is invalid", http.StatusBadRequest)
		return
	}

	featureTree, featureIds, err := s.featureStore.FindRelationsByID(ctx, featureId)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}

	_, err = s.listingStore.FindByListingIDs(ctx, featureIds)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}

	response, err := json.Marshal(featureTree)
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
