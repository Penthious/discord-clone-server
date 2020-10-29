package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201027180933, Down20201027180933)
}

func Up20201027180933(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func Down20201027180933(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
