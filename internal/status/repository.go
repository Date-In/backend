package status

import "dating_service/pkg/db"

type StatusRepository struct {
	db *db.Db
}

func NewStatusRepository(db *db.Db) *StatusRepository {
	return &StatusRepository{db}
}
