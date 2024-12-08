package model

import (
	"github.com/google/uuid"
	"time"
)

type ProjectRequest struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

type Project struct {
	Guid      uuid.UUID  `json:"guid"`
	Alias     string     `json:"alias"`
	Name      string     `json:"name"`
	Info      *string    `json:"info"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
