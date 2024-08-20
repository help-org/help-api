package database

import (
	"context"
	"log"

	db "directory/pkg/database"
	"directory/pkg/types"
)

type DivisionStore struct {
	store db.Pool
}

func NewDivisionStore(s db.Pool) *DivisionStore {
	return &DivisionStore{store: s}
}

func (s *DivisionStore) Create(ctx context.Context, division types.Division) (createdId int, err error) {
	err = s.store.QueryRow(ctx, createQuery, division.Name, division.Type, division.ParentId).Scan(&createdId)
	return
}

func (s *DivisionStore) FindByID(ctx context.Context, id int) (division *types.Division, err error) {
	division = &types.Division{}
	err = s.store.QueryRow(ctx, findByIDQuery, id).Scan(&division.Id, &division.Name, &division.Type, &division.ParentId)
	return
}

func (s *DivisionStore) Update(ctx context.Context, division types.Division) (updatedId int, err error) {
	err = s.store.QueryRow(ctx, updateQuery, division.Id, division.Name, division.Type, division.ParentId).Scan(&updatedId)
	return
}

func (s *DivisionStore) Delete(ctx context.Context, id int) (deletedId int, err error) {
	err = s.store.QueryRow(ctx, deleteQuery, id).Scan(&deletedId)
	return
}

func (s *DivisionStore) FindRelationsByID(ctx context.Context, id int) (divisions []*types.Division, err error) {
	rows, err := s.store.Query(ctx, recursiveFindByIDQuery, id)
	if err != nil {
		log.Fatalf("failed to query row: %v", err)
	}

	for rows.Next() {
		var d types.Division
		err := rows.Scan(&d.Id, &d.Name, &d.Type, &d.ParentId)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		divisions = append(divisions, &d)
	}

	return
}

const createQuery = "INSERT INTO directory.divisions (name, type, parent_id) VALUES ($1, $2, $3) RETURNING id"

const findByIDQuery = "SELECT d.id, d.name, d.type, d.parent_id FROM directory.divisions d WHERE id = $1"

const updateQuery = "UPDATE directory.divisions SET name = $2, type = $3, parent_id = $4 WHERE id = $1 RETURNING id"

const deleteQuery = "DELETE FROM directory.divisions WHERE id = $1 RETURNING id"

// WARNING Can this ever be circular?
const recursiveFindByIDQuery = `
	WITH RECURSIVE ParentCTE AS (
		-- Start with the given record and find its children
		SELECT id, name, type, parent_id
		FROM directory.divisions
		WHERE id = $1
		
		UNION ALL
		
		-- Find all parents of the current record
		SELECT loc.id, loc.name, loc.type, loc.parent_id
		FROM directory.divisions loc
		JOIN ParentCTE p ON loc.id = p.parent_id
	),
	ChildCTE AS (
		-- Start with the given record and find its children
		SELECT id, name, type, parent_id
		FROM directory.divisions
		WHERE id = $1
		
		UNION ALL
		
		-- Find all children of the current record
		SELECT loc.id, loc.name, loc.type, loc.parent_id
		FROM directory.divisions loc
		JOIN ChildCTE c ON loc.parent_id = c.id
	)
	-- Combine results from both ParentCTE and ChildCTE
	SELECT * FROM ParentCTE
	UNION
	SELECT * FROM ChildCTE
	ORDER BY id;
`
