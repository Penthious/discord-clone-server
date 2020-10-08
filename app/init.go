package app

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	DB *gorm.DB
}

func InitServices() (Server, error) {
	var s Server
	var err error

	s.DB, err = InitDB(s)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v \n", err.Error())
	}

	return s, err
}

func InitDB(s Server) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	connectionInfo := "root:secret@tcp(db:3306)/discord_clone?parseTime=true"
	// return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=%s", c.User, c.Password, c.Protocol, c.Host, c.Port, c.Name, c.ParseTime)

	return gorm.Open(mysql.Open(connectionInfo), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
}
