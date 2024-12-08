package storage

import (
	"afera-projects/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config *config.Config
	db     *sql.DB
}

func NewStorage(config *config.Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) GetDatabaseURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host, s.config.Port, s.config.User, s.config.Pass, s.config.Name, s.config.SSLMode)
}

func (s *Storage) GetMigrationURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		s.config.User, s.config.Pass, s.config.Host, s.config.Port, s.config.Name, s.config.SSLMode)
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.GetDatabaseURI())
	if err != nil {
		return fmt.Errorf("could not open connection to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}

	s.db = db

	log.Println("DB connection...")

	return nil
}

func (s *Storage) Close() error {
	if s.db != nil {

		return s.db.Close()
	}

	return nil
}

func (s *Storage) GetDB() *sql.DB {
	return s.db
}
