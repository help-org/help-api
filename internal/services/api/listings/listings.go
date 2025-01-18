package api

import (
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

func (s *ListingService) RegisterRoutes(mux *chi.Mux) {}
