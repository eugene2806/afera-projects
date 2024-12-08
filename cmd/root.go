package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"my-template/internal/config"
)

const migratePath = "file://migrate"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-template",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cfg := config.NewConfig()
	cfg.ConfigStorageField()
	migrate := migrateCmd()

	migrate.AddCommand(migrateUpCmd(cfg.Storage.GetMigrationURL(), migratePath))
	migrate.AddCommand(migrateDownCmd(cfg.Storage.GetMigrationURL(), migratePath))

	rootCmd.AddCommand(restCmd())
	rootCmd.AddCommand(migrate)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {

}
