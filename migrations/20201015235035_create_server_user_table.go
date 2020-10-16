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
CREATE TABLE server_users (
  user_id int(11) NOT NULL,
  server_id int(11) NOT NULL,
  created_at datetime,
  deleted_at datetime,
  updated_at datetime,
  PRIMARY KEY (user_id, server_id),
  CONSTRAINT FK_servers_TO_server_users FOREIGN KEY (server_id) REFERENCES servers (id),
  CONSTRAINT FK_users_TO_server_users FOREIGN KEY (user_id) REFERENCES users (id)
)


ENGINE = 'InnoDB'
DEFAULT CHARSET = 'utf8mb4'
COLLATE = 'utf8mb4_unicode_ci'
`)

	/*
		CREATE TABLE `server_users` (
		  `server_id` bigint unsigned,
		  `user_id` bigint unsigned,
		  PRIMARY KEY (`server_id`,`user_id`),
		  CONSTRAINT `fk_server_users_server` FOREIGN KEY (`server_id`) REFERENCES `servers`(`id`),
		  CONSTRAINT `fk_server_users_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
		)
	*/
	if err != nil {
		return err
	}

	return nil
}

func Down20201015235002(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE server_users`)
	if err != nil {
		return err
	}
	return nil
}
