package utils

import (
	"discord-clone-server/models"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "discord-clone-server/migrations"

	"github.com/pressly/goose"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTestDB(t *testing.T, dbName string) *gorm.DB {
	var dsn string

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	dsn = GetMysqlDSN("")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	fmt.Printf("db: %v", db.Name())
	return db
}

func DropTestDB(t *testing.T, db *gorm.DB, dbName string) {
	if err := db.Exec(fmt.Sprintf("DROP DATABASE if exists `%s`", dbName)).Error; err != nil {
		t.Fatalf("error refreshing test DB: %s", err.Error())
	}
}

func MakeTestUsers(t *testing.T, db *gorm.DB, users []models.User) {

	tx := db.Create(&users)

	if err := tx.Error; err != nil {
		t.Fatalf("error creating users: %s", err.Error())
	}
}

func NewTestRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	// return redismock.NewNiceMock(client)
}
