package api

import (
	"directory/internal/store/database"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ListingService struct {
	listingStore database.ListingStore
	contactStore database.ContactsStore
}

func NewListingService(listingStore database.ListingStore, contactStore database.ContactsStore) *ListingService {
	return &ListingService{
		contactStore: contactStore,
		listingStore: listingStore,
	}
}

func (s *ListingService) RegisterRoutes(mux *chi.Mux) {
	mux.Get("/listing/{id}/contacts", s.FindContactByListingID)
}

func (s *ListingService) FindContactByListingID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	contacts, err := s.contactStore.FindContactsByListingIDs(ctx, id)
	if err != nil {
		http.Error(w, "contacts were not found for the requested listing", http.StatusNotFound)
	}

	response, err := json.Marshal(contacts)
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
