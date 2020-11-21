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
	_, err := tx.Exec(`
CREATE TABLE categories (
  id        int(11)      NOT NULL AUTO_INCREMENT,
  name      varchar(100) NOT NULL,
  server_id int(11)      NOT NULL,
  
  PRIMARY KEY (id, server_id),

  FOREIGN KEY (server_id) REFERENCES servers (id)
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

func Down20201027180933(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE categories`)
	if err != nil {
		return err
	}
	return nil
}
