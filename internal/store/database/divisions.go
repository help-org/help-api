package database

import (
	"context"
	"fmt"

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

func (s *DivisionStore) Delete(ctx context.Context, id int) (err error) {
	execResult, err := s.store.Exec(ctx, deleteQuery, id)
	if err != nil {
		return err
	}

	if execResult.RowsAffected() == 0 {
		return fmt.Errorf("no division found with id %d", id)
	}
	return
}

func (s *DivisionStore) Update(ctx context.Context, division types.Division) (id int, err error) {
	err = s.store.QueryRow(ctx, updateQuery, division.Name, division.Type, division.ParentId).Scan(&id)
	return
}

const createQuery = "INSERT INTO directory.divisions (name, type, parent_id) VALUES ($1, $2, $3) RETURNING id"

const findByIDQuery = "SELECT d.id, d.name, d.type, d.parent_id FROM directory.divisions d WHERE id = $1"

const deleteQuery = "DELETE FROM divisions WHERE id = $1"

const updateQuery = "UPDATE directory.divisions SET name = $1, type = $2, parent_id = $3 WHERE id = $4 RETURNING id"
