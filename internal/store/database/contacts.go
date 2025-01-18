package database

import (
	"context"
	"fmt"

	db "directory/pkg/database"
	"directory/pkg/types"
)

type ContactsStore struct {
	store db.Pool
}

func NewContactsStore(s db.Pool) *ContactsStore {
	return &ContactsStore{store: s}
}

func (s *ContactsStore) FindContactsByListingIDs(ctx context.Context, listingId string) (contacts []*types.Contact, err error) {
	rows, err := s.store.Query(ctx, findContactsByIDs, listingId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	for rows.Next() {
		var contact types.Contact
		if err := rows.Scan(&contact.Id, &contact.Name, &contact.Type, &contact.Details); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		contacts = append(contacts, &contact)
	}
	return
}

const findContactsByIDs = `
	SELECT
		id,
		name,
		type,
		details
	FROM contacts
	WHERE internal_id
    	= ANY(SELECT unnest(contact_ids) FROM listings WHERE id = $1)
`

// Notes
//LEFT JOIN LATERAL
//unnest(l.contact_ids) AS contact_id ON true
//LEFT JOIN
//contacts c ON c.internal_id = contact_id
