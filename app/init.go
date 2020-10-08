package app

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitServices() {
	InitDB()
}

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Error,
			Colorful: true,
		},
	)

	connectionInfo := "mysql:secret@tcp(127.0.0.1:3306)/test?parseTime=true"
	// return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=%s", c.User, c.Password, c.Protocol, c.Host, c.Port, c.Name, c.ParseTime)

	_, err := gorm.Open(mysql.Open(connectionInfo), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("error opening gorm db: %v \n", err.Error())
		// return errors.Wrap(err, "error connecting to DB")
	}

	// s.DB = db
	// return nil
}
