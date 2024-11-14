package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"my-template/internal/builder"
	"my-template/internal/transport/server"
)

func restCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "rest",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			serv := server.BuildServer()
			handler := builder.NewHandlerBuilder(serv).BuildHandler()
			restServer := builder.BuildRestServer("8080", handler)
			log.Println("rest server start")
			log.Fatal(restServer.ListenAndServe())
		},
	}
}
