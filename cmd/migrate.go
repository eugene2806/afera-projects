package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
)

func migrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Manage database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("I'm migrate")
		},
	}
}

func migrateUpCmd(MigrationURL, dbPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dbURL, err := migrate.New(dbPath, MigrationURL)

			if err != nil {
				log.Fatalf("Failed to create migrate instance: %v", err)
			}

			if err = dbURL.Up(); err != nil {
				log.Fatalf("Failed to apply migrations: %v", err)
			}

			log.Println("Successfully applied up migrations")
		},
	}
}

func migrateDownCmd(MigrationURL, dbPath string) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Apply all down migrations",
		Run: func(cmd *cobra.Command, args []string) {
			dbURL, err := migrate.New(dbPath, MigrationURL)

			if err != nil {
				log.Fatalf("Failed to create migrate instance: %v", err)
			}

			if err = dbURL.Down(); err != nil {
				log.Fatalf("Failed to apply migrations: %v", err)
			}

			log.Println("Successfully applied down migrations")
		},
	}
}
