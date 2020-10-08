package app

import (
	"discord-clone-server/repositories"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Services struct {
	DB       *gorm.DB
	UserRepo repositories.UserRepo
}

func InitServices() (Services, error) {
	var s Services
	var err error

	s.DB, err = InitDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v \n", err.Error())
	}

	s.UserRepo = InitUserRepo(s.DB)

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

	connectionInfo := "root:secret@tcp(172.21.0.2:3306)/discord_clone?parseTime=true"
	// return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=%s", c.User, c.Password, c.Protocol, c.Host, c.Port, c.Name, c.ParseTime)

	return gorm.Open(mysql.Open(connectionInfo), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
}

func InitUserRepo(db *gorm.DB) repositories.UserRepo {
	return repositories.NewUserRepo(db)
}
