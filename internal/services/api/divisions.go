package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"directory/internal/store/database"
	"directory/pkg/types"
	"directory/pkg/logger"
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
	mux.Post("/divisions", s.Create)
}

func (s *DivisionService) Create(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    logger.Info(ctx, "Create Error")
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        http.Error(w, "Invalid JSON request", http.StatusBadRequest)
        return
    }

    name := data["name"].(string)
    dtype := data["type"].(string)
    parentId := int(data["parentId"].(float64))

    division := types.Division{
        Name: name,
        Type: dtype,
        ParentId: &parentId,
    }

    id, err := s.divisionStore.Create(ctx, division)

    if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
        return
    }

    logger.Info(ctx, "Error with response")

    response, err := json.Marshal(id)
    if err != nil {
        panic(err) // Handle error appropriately
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
