package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"directory/internal/store/database"
	"directory/pkg/types"
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
	mux.Post("/divisions", s.Create)
	mux.Get("/divisions/{id}", s.FindByID)
	mux.Put("/divisions/{id}", s.Update)
	mux.Delete("/divisions/{id}", s.Delete)
}

func (s *DivisionService) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to unmarshal body as JSON", http.StatusBadRequest)
		return
	}

	name := data["name"].(string)
	divisionType := data["type"].(string)
	parentId := int(data["parent_id"].(float64))

	division := types.Division{
		Name:     name,
		Type:     divisionType,
		ParentId: &parentId,
	}

	id, err := s.divisionStore.Create(ctx, division)
	if err != nil {
		http.Error(w, "Failed to create division", http.StatusInternalServerError)
		return
	}
	division.Id = id

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

func (s *DivisionService) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "division id is invalid", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to update request body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to unmarshal body as JSON", http.StatusBadRequest)
		return
	}

	division, err := s.divisionStore.FindByID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to find division", http.StatusNotFound)
		return
	}

	if data["name"] != nil {
		division.Name = data["name"].(string)
	}
	if data["type"] != nil {
		division.Type = data["type"].(string)
	}
	if data["parent_id"] != nil {
		parentId := int(data["parent_id"].(float64))
		division.ParentId = &parentId
	}

	_, err = s.divisionStore.Update(ctx, *division)
	if err != nil {
		http.Error(w, "Failed to update division", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(division)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)

	if err != nil {
		http.Error(w, "Error updating division", http.StatusInternalServerError)
		return
	}
}

func (s *DivisionService) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Division ID", http.StatusBadRequest)
		return
	}

	_, err = s.divisionStore.Delete(ctx, id)
	if err != nil {
		http.Error(w, "Cannot delete division", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
