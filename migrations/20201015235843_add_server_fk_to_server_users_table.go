package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201015235843, Down20201015235843)
}

func Up20201015235843(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
ALTER TABLE server_user
ADD CONSTRAINT FK_servers_TO_servers_users
FOREIGN KEY (server_id)
REFERENCES servers (id)
`)
	if err != nil {
		return err
	}

	return nil
}

func Down20201015235843(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
ALTER TABLE server_user
DROP FOREIGN KEY (server_id)
`)
	if err != nil {
		return err
	}

	return nil
}
