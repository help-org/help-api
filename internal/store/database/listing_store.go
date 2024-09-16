package database

import (
	"context"

	db "directory/pkg/database"
	"directory/pkg/types"
)

type ListingStore struct {
	store db.Pool
}

func NewListingStore(s db.Pool) *ListingStore {
	return &ListingStore{store: s}
}

func (s *ListingStore) FindByListingIDs(ctx context.Context, featureIds []int) (listings *[]types.Listing, err error) {
	err = s.store.QueryRow(ctx, findContactsByListingIDs, featureIds).Scan(&listings.Id, &listings.Name, &listings.Type, &listings.ParentId)
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
		contacts c ON c.internal_id = contact_id;
	WHERE l.internal_id in ($1)
`
