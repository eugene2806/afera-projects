package repository

import (
	"afera-projects/internal/model"
	"afera-projects/storage"
	"database/sql"
	"github.com/google/uuid"
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

func (p *ProjectRepository) GetByID(id uuid.UUID) (interface{}, error) {
	project := model.Project{}

	query := "SELECT guid, name, info, created_at, updated_at FROM projects WHERE guid = $1"

	err := p.db.QueryRow(query, id).Scan(&project.Guid, &project.Name, &project.Info, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {

		return nil, err
	}

	response := struct {
		Guid      uuid.UUID `json:"guid"`
		Name      string    `json:"name"`
		Info      *string   `json:"info"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		Guid:      project.Guid,
		Name:      project.Name,
		Info:      project.Info,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}

	return response, nil
}

func (p *ProjectRepository) Create(req model.ProjectRequest) () {

}
