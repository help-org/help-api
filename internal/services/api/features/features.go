package api

import (
	"directory/pkg/types"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"directory/internal/store/database"
)

type FeatureService struct {
	featureStore database.FeatureStore
	listingStore database.ListingStore
}

func NewFeatureService(featuresStore database.FeatureStore, listingStore database.ListingStore) *FeatureService {
	return &FeatureService{
		featureStore: featuresStore,
		listingStore: listingStore,
	}
}

func (s *FeatureService) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/features/{id}", s.FindByID)
	mux.Get("/features/{featureId}/listings", s.FindListingsByFeatureId)

}

func (s *FeatureService) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	features, err := s.featureStore.FindByID(ctx, id)
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

func (s *FeatureService) FindListingsByFeatureId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	featureId := chi.URLParam(r, "featureId")

	featureTree, featureInternalIds, err := s.featureStore.FindRelationsByID(ctx, featureId)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}
	listings, err := s.listingStore.FindByListingFeatureInternalIDs(ctx, featureInternalIds)
	if err != nil {
		http.Error(w, "listings for feature ids were not found", http.StatusNotFound)
	}

	listingMap := listingsToMap(listings)
	for _, feature := range featureTree {
		listing, match := listingMap[feature.InternalId]
		if match {
			feature.Listings = append(feature.Listings, listing)
		}
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

func listingsToMap(listings []*types.Listing) map[int]*types.Listing {
	i := map[int]*types.Listing{}
	for _, listing := range listings {
		i[listing.FeatureId] = listing
	}
	return i
}
