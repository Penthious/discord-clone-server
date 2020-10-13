package repositories

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/pressly/goose"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTestDB(t *testing.T, dbName string) *gorm.DB {
	// var dsn string

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	connectionInfo := "root:secret@tcp(172.21.0.2:3306)/?parseTime=true"

	db, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		t.Fatalf("error opening gorm connection: %v\n", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("error opening gorm connection: %v\n", err.Error())
	}

	sqlDB.SetConnMaxLifetime(30 * time.Second)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	_, err = sqlDB.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		t.Fatalf("error connecting to DB: %v\n", err.Error())
	}

	_, err = sqlDB.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		t.Fatalf("error connecting to DB: %v\n", err.Error())
	}

	_, err = sqlDB.Exec("USE " + dbName)
	if err != nil {
		t.Fatalf("error connecting to DB: %v\n", err.Error())
	}

	// run migrations
	if err := goose.SetDialect("mysql"); err != nil {
		t.Fatalf("error setting dialect: %s", err.Error())
	}
	if err := goose.Run("up", sqlDB, "."); err != nil {
		t.Fatalf("error running migrations: %v\n", err.Error())
	}

	return db
}

func GetDBName(prefix string) string {
	length := 6
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(charset[rand.Intn(len(charset))])
	}

	return fmt.Sprintf("%s_%s", prefix, string(runes))
}

func DropTestDB(t *testing.T, db *gorm.DB, dbName string) {
	if err := db.Exec(fmt.Sprintf("DROP DATABASE if exists `%s`", dbName)).Error; err != nil {
		t.Fatalf("error refreshing test DB: %s", err.Error())
	}
}
