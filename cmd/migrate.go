package cmd

import (
	"discord-clone-server/app"

	_ "discord-clone-server/migrations"

	"github.com/pkg/errors"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

// migrate command definition
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "goose migrations (go run main.go migrate up)",
	RunE:  migrate,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	MigrateCmd.SetUsageFunc(func(c *cobra.Command) error {
		c.Println(`
Usage: address-qual migrate [OPTIONS] COMMAND

Drivers:
    postgres
    mysql
    sqlite3
    redshift

Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp

Examples:
    discord-clone migrate status
    discord-clone migrate create init sql
    discord-clone migrate create add_some_column sql
    discord-clone migrate create fetch_user_data go
    discord-clone migrate up

    discord-clone migrate status`)
		return nil
	})
}

func migrate(_ *cobra.Command, args []string) error {
	db, err := app.InitDB()
	if err != nil {
		return errors.Wrap(err, "error initializing migration connection")
	}

	var arguments []string
	if len(args) > 1 {
		arguments = args[1:]
	}

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return goose.Run(args[0], sqlDB, ".", arguments...)
}
