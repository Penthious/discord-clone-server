package repositories

import (
	"discord-clone-server/models"
	"discord-clone-server/utils"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "discord-clone-server/migrations"

	"github.com/pressly/goose"

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

	dsn = utils.GetMysqlDSN("")

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

func MakeTestUsers(t *testing.T, db *gorm.DB, users []models.User) {

	tx := db.Create(&users)

	if err := tx.Error; err != nil {
		t.Fatalf("error creating users: %s", err.Error())
	}
}
