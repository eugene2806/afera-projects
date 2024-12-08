package repository

import (
	"database/sql"
	"my-template/storage"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(storage *storage.Storage) *ProjectRepository {
	return &ProjectRepository{
		db: storage.GetDB(),
	}
}
