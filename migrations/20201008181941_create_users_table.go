package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201008181941, Down20201008181941)
}

func Up20201008181941(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  first_name varchar(100)   ,
  last_name  varchar(100)   ,
  email      varchar(100)   UNIQUE,
  username   varchar(100)   UNIQUE,
  password   varchar(100)  ,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime NOT NULL,
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

func Down20201008181941(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE users`)
	if err != nil {
		return err
	}
	return nil
}
