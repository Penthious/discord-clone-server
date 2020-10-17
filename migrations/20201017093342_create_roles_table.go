package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201017093342, Down20201017093342)
}

func Up20201017093342(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE roles (
  id         int(11) NOT NULL AUTO_INCREMENT,
  name       varchar(100),
  server_id  int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  PRIMARY KEY (id, server_id),
  CONSTRAINT FK_servers_TO_server_users FOREIGN KEY (server_id) REFERENCES servers (id)
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

func Down20201017093342(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE roles`)
	if err != nil {
		return err
	}
	return nil
	return nil
}
