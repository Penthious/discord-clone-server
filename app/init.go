package app

import (
	"discord-clone-server/repositories"
	"discord-clone-server/utils"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Services struct {
	DB         *gorm.DB
	UserRepo   repositories.UserRepo
	ServerRepo repositories.ServerRepo
}

func InitServices() (Services, error) {
	var s Services
	var err error

	s.DB, err = InitDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v \n", err.Error())
	}

	s.UserRepo = InitUserRepo(s.DB)
	s.ServerRepo = InitServerRepo(s.DB)

	return s, err
}

func InitDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	dsn := utils.GetMysqlDSN("")

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
}

func InitUserRepo(db *gorm.DB) repositories.UserRepo {
	return repositories.NewUserRepo(db)
}
func InitServerRepo(db *gorm.DB) repositories.ServerRepo {
	return repositories.NewServerRepo(db)
}
