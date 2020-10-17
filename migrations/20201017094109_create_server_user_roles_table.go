package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201017094109, Down20201017094109)
}

func Up20201017094109(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE server_user_roles (
  id         int(11) NOT NULL,
  server_id  int(11) NOT NULL,
  user_id    int(11) NOT NULL,
  role_id    int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  PRIMARY KEY (id, server_id, user_id, role_id),
  CONSTRAINT FK_servers_TO_server_user_roles FOREIGN KEY (server_id) REFERENCES servers (id),
  CONSTRAINT FK_users_TO_server_user_roles FOREIGN KEY (user_id) REFERENCES users (id),
  CONSTRAINT FK_roles_TO_server_user_roles FOREIGN KEY (role_id) REFERENCES roles (id)
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

func Down20201017094109(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE server_user_roles`)
	if err != nil {
		return err
	}
	return nil
}
