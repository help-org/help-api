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

const createQuery = "INSERT INTO directory.divisions (name, type, parent_id) VALUES ($1, $2, $3) RETURNING id"

const findByIDQuery = "SELECT d.id, d.name, d.type, d.parent_id FROM directory.divisions d WHERE id = $1"

const updateQuery = "UPDATE directory.divisions SET name = $2, type = $3, parent_id = $4 WHERE id = $1 RETURNING id"

const deleteQuery = "DELETE FROM directory.divisions WHERE id = $1 RETURNING id"
