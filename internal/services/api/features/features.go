package api

import (
	"directory/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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

// TODO add query param for find 1 or recursive
func (s *FeatureService) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "feature id is invalid", http.StatusBadRequest)
		return
	}

	features, _, err := s.featureStore.FindRelationsByID(ctx, id)
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

// TODO add query param for find 1 or recursive
func (s *FeatureService) FindListingsByFeatureId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	featureId, err := strconv.Atoi(chi.URLParam(r, "featureId"))
	if err != nil {
		http.Error(w, "feature id is invalid", http.StatusBadRequest)
		return
	}

	featureTree, featureInternalIds, err := s.featureStore.FindRelationsByID(ctx, featureId)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}
	listings, err := s.listingStore.FindByListingFeatureInternalIDs(ctx, featureInternalIds)
	listingMap := listingsToMap(listings)
	if err != nil {
		http.Error(w, "feature id was not found", http.StatusNotFound)
	}
	fmt.Println(listingMap)
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
