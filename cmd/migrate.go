package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func migrateCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "migrate",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("I'm migrate")
		},
	}
}
