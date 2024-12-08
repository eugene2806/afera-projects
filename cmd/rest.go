package cmd

import (
	"afera-projects/internal/builder"
	"afera-projects/internal/repository"
	"afera-projects/internal/transport/server"
	"afera-projects/storage"
	"github.com/spf13/cobra"
	"log"
)

func restCmd(stor *storage.Storage) *cobra.Command {

	return &cobra.Command{
		Use:   "rest",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			reposit := repository.NewProjectRepository(stor)

			serv := server.BuildServer(reposit)

			handler := builder.NewHandlerBuilder(serv).BuildHandler()

			restServer := builder.BuildRestServer("50051", handler)

			log.Println("rest server start...")

			log.Fatal(restServer.ListenAndServe())
		},
	}
}
