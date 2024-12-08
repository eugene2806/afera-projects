package repository

import (
	"afera-projects/internal/errors_pkg"
	"afera-projects/internal/model"
	"afera-projects/storage"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"strconv"
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

func (p *ProjectRepository) GetAllProjects(pageStr, limitStr string) ([]*model.Project, int, int, error) {
	var fullCount, fullPage int

	page, err := strconv.Atoi(pageStr)
	if err != nil {

		return nil, 0, 0, err
	}

	if page <= 0 {

		return nil, 0, 0, fmt.Errorf("page is %w", errors_pkg.ErrLessZero)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {

		return nil, 0, 0, err
	}

	if limit <= 0 {

		return nil, 0, 0, fmt.Errorf("limit is %w", errors_pkg.ErrLessZero)
	}

	offset := (page - 1) * limit

	query1 := `SELECT
				COUNT(*) AS full_count,
				CEIL(COUNT(*)::DECIMAL / $1) AS full_page
				FROM projects;`

	err = p.db.QueryRow(query1, limit).Scan(&fullCount, &fullPage)

	if err != nil {

		return nil, 0, 0, err
	}

	query2 := `SELECT guid, alias, name, info, created_at, updated_at
			FROM projects
			LIMIT $1 OFFSET $2`

	rows, err := p.db.Query(query2, limit, offset)

	if err != nil {

		return nil, 0, 0, err
	}

	defer rows.Close()

	var projects []*model.Project

	for rows.Next() {
		var project model.Project

		err = rows.Scan(&project.Guid, &project.Alias, &project.Name, &project.Info, &project.CreatedAt, &project.UpdatedAt)

		if err != nil {

			return nil, 0, 0, err

		}

		projects = append(projects, &project)
	}

	if err = rows.Err(); err != nil {

		return nil, 0, 0, err
	}

	return projects, fullCount, fullPage, nil
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

func (p *ProjectRepository) Create(req model.ProjectRequest) {

}
