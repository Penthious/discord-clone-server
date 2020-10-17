package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201017113733, Down20201017113733)
}

func Up20201017113733(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE role_permissions (
  id         int(11) NOT NULL AUTO_INCREMENT,
  role_id    int(11) NOT NULL,
  permission_id  int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,

  PRIMARY KEY (id, role_id, permission_id),

  FOREIGN KEY (role_id) REFERENCES roles (id),
  FOREIGN KEY (permission_id) REFERENCES permissions (id)
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

func Down20201017113733(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE role_permissions`)
	if err != nil {
		return err
	}
	return nil
}
