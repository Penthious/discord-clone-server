package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201015235024, Down20201015235024)
}

func Up20201015235024(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
CREATE TABLE servers (
  id         int(11) NOT NULL AUTO_INCREMENT,
  name       varchar(100),
  private    boolean NOT NULL DEFAULT 0,
  user_id    int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  
  PRIMARY KEY (id)
)
ENGINE = 'InnoDB'
DEFAULT CHARSET = 'utf8mb4'
COLLATE = 'utf8mb4_unicode_ci'
`)
	// CREATE TABLE `servers` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`name` longtext,PRIMARY KEY (`id`),INDEX idx_servers_deleted_at (`deleted_at`))

	if err != nil {
		return err
	}

	return nil
}

func Down20201015235024(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE servers`)
	if err != nil {
		return err
	}
	return nil
}
