package cmd

import (
	"discord-clone-server/app"
	"log"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run the web server",
	RunE:  server,
	Args:  cobra.ExactArgs(0),
}

func server(_ *cobra.Command, _ []string) error {
	_, err := app.InitServices()
	if err != nil {
		log.Fatalf("error initializing services: %v\n", err.Error())
	}
	return nil
}
