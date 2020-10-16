package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201015235002, Down20201015235002)
}

func Up20201015235002(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE server_user (
  user_id int(11) NOT NULL,
  server_id int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime NOT NULL,
  PRIMARY KEY (user_id, server_id)
)
ENGINE = 'InnoDB'
DEFAULT CHARSET = 'utf8mb4'
COLLATE = 'utf8mb4_unicode_ci'
`)
	if err != nil {
		return err
	}

	return nil
}

func Down20201015235002(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE server_user`)
	if err != nil {
		return err
	}
	return nil
}
