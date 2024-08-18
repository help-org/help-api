package database

import (
	"context"

	db "directory/pkg/database"
	"directory/pkg/types"
)

type DivisionStore struct {
	store db.Pool
}

func NewDivisionStore(s db.Pool) *DivisionStore {
	return &DivisionStore{store: s}
}

func (s *DivisionStore) FindByID(ctx context.Context, id int) (division *types.Division, err error) {
	division = &types.Division{}
	err = s.store.QueryRow(ctx, findByIDQuery, id).Scan(&division.Id, &division.Name, &division.Type, &division.ParentId)
	return
}

func (s *DivisionStore) Create(ctx context.Context, division types.Division) (id int, err error) {
	err = s.store.QueryRow(ctx, createQuery, division.Name, division.Type, division.ParentId).Scan(&id)
	return
}

const createQuery = "INSERT INTO directory.divisions (name, type, parent_id) VALUES ($1, $2, $3) RETURNING id"

const findByIDQuery = "SELECT d.id, d.name, d.type, d.parent_id FROM directory.divisions d WHERE id = $1"
