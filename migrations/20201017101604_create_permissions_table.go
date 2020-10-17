package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201017101604, Down20201017101604)
}

func Up20201017101604(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE permissions (
  id         int(11)      NOT NULL AUTO_INCREMENT,
  permission varchar(100) UNIQUE,
  name       varchar(100) UNIQUE,
  detail     varchar(100),
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  PRIMARY KEY (id)
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

func Down20201017101604(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE permissions`)
	if err != nil {
		return err
	}
	return nil
}
