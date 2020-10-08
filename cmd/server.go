package cmd

import (
	"discord-clone-server/app"
	"fmt"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run the web server",
	RunE:  Server,
	Args:  cobra.ExactArgs(0),
}

func Server(_ *cobra.Command, _ []string) error {
	s, err := app.InitServices()
	if err != nil {
		return err
	}
	fmt.Print(s)

	defer s.DB.Close()
	return nil
}
