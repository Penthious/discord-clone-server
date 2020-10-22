package cmd

import (
	"discord-clone-server/app"
	"discord-clone-server/seeder"

	_ "discord-clone-server/migrations"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// migrate command definition
var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the DB",
	RunE:  seed,
	Args:  cobra.MinimumNArgs(0),
}

func init() {
	MigrateCmd.SetUsageFunc(func(c *cobra.Command) error {
		c.Println("Usage:  discord-clone seed")
		return nil
	})
}

func seed(_ *cobra.Command, args []string) error {
	db, err := app.InitDB()
	if err != nil {
		return errors.Wrap(err, "error initializing migration connection")
	}

	// var arguments []string
	// if len(args) > 1 {
	// 	arguments = args[1:]
	// }

	if err := seeder.PermissionsSeeder(db); err != nil {
		return err
	}

	return err
}
