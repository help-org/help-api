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

func (s *ListingStore) FindByListingFeatureInternalIDs(ctx context.Context, featureIds []int) (listings []*types.Listing, err error) {
	rows, err := s.store.Query(ctx, findContactsByListingIDs, featureIds)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	for rows.Next() {
		var listing types.Listing
		if err := rows.Scan(&listing.Id, &listing.Name, &listing.Type, &listing.FeatureId, &listing.Address, &listing.Details, &listing.ContactIds); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		listings = append(listings, &listing)
	}
	return
}

const findContactsByListingIDs = `
	SELECT
		l.id,
		l.name,
		l.type,
		l.feature_id,
		l.address,
		l.details,
		l.contact_ids
	FROM
		listings l
	WHERE l.feature_id = ANY($1)
`
