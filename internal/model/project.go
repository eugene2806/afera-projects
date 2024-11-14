package model

import "time"

type Project struct {
	Guid      string
	Alias     string
	Name      string
	Desc      *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
