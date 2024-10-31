package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func restCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "rest",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("I'm rest")
		},
	}
}
