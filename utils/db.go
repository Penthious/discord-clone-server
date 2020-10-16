package utils

import (
	"fmt"
	"os"
	"time"

	mySQL "github.com/go-sql-driver/mysql"
)

func GetMysqlDSN(prefix string) string {
	config := &mySQL.Config{
		User:                 os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_USER")),
		Passwd:               os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_PASSWORD")),
		Addr:                 os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_ADDR")),
		DBName:               os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_NAME")),
		Net:                  "tcp",
		Timeout:              5 * time.Second,
		ReadTimeout:          5 * time.Second,
		WriteTimeout:         5 * time.Second,
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return config.FormatDSN()
}
