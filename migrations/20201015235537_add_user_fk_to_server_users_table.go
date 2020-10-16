package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201015235537, Down20201015235537)
}

func Up20201015235537(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
ALTER TABLE server_user
ADD CONSTRAINT FK_users_TO_servers_users
FOREIGN KEY (user_id)
REFERENCES users (id)
`)
	if err != nil {
		return err
	}

	return nil
}

func Down20201015235537(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
ALTER TABLE server_user
DROP FOREIGN KEY (user_id)
`)
	if err != nil {
		return err
	}

	return nil
}
