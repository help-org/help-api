package directory

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Service struct {
	database *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{database: db}
}

func (s *Service) ListLocal(writer http.ResponseWriter, request *http.Request) {
	directory := &Directory{
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

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		return
	}
}
