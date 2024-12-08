package repository

import (
	"afera-projects/internal/errors_pkg"
	"afera-projects/internal/model"
	"afera-projects/storage"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"time"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(storage *storage.Storage) *ProjectRepository {
	return &ProjectRepository{
		db: storage.GetDB(),
	}
}

func (p *ProjectRepository) Create(req model.ProjectRequest) (interface{}, error) {
	if req.Name == "" || req.Info == "" {

		return nil, fmt.Errorf("name or info %w", errors_pkg.ErrInvalidRequest)
	}

	var id = uuid.New()
	var alias = slug.Make(req.Name)
	var createdAt time.Time // Время мы указываем сами или оно выставляется в базе?

	query := `INSERT INTO projects (guid, alias, name, info)
					VALUES ($1, $2, $3, $4)
					RETURNING created_at`

	err := p.db.QueryRow(query, id, req.Name, req.Info, alias).Scan(&createdAt)
	if err != nil {

		return nil, err
	}

	response := struct {
		Guid      uuid.UUID `json:"guid"`
		Name      string    `json:"name"`
		Info      string    `json:"info"`
		CreatedAt time.Time `json:"created_at"`
		Alias     string    `json:"alias"`
	}{
		Guid:      id,
		Name:      req.Name,
		Info:      req.Info,
		CreatedAt: createdAt,
		Alias:     alias,
	}

	return response, nil

}
