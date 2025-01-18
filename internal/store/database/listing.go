package database

import (
	"context"
	"fmt"

	db "directory/pkg/database"
	"directory/pkg/types"
)

type ListingStore struct {
	store db.Pool
}

func NewListingStore(s db.Pool) *ListingStore {
	return &ListingStore{store: s}
}

func (s *ListingStore) FindByListingIDs(ctx context.Context, featureIds []*int) (listings []*types.Listing, err error) {
	rows, err := s.store.Query(ctx, findContactsByListingIDs, featureIds)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	// Iterate over the rows and map to struct
	for rows.Next() {
		var listing types.Listing
		if err := rows.Scan(&listing.Id, &listing.Name, &listing.Type, &listing.ParentId); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		listings = append(listings, &listing)
	}
	return
}

const findContactsByListingIDs = `
	SELECT
		l.id AS listing_id,
		l.name AS listing_name,
		l.details AS listing_details,
		l.contact_ids,                    -- Array of contact_ids
		c.id AS contact_id,
		c.name AS contact_name,
		c.type AS contact_type,
		c.details AS contact_details
	FROM
		listings l
	LEFT JOIN LATERAL
		unnest(l.contact_ids) AS contact_id ON true
	LEFT JOIN
		contacts c ON c.internal_id = contact_id
	WHERE l.internal_id in ($1)
`
