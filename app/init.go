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
	DB             *gorm.DB
	UserRepo       repositories.UserRepo
	ServerRepo     repositories.ServerRepo
	RoleRepo       repositories.RoleRepo
	PermissionRepo repositories.PermissionRepo
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
	s.RoleRepo = InitRoleRepo(s.DB)
	s.PermissionRepo = InitPermissionRepo(s.DB)

	return s, err
}

func InitDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			// LogLevel: logger.Error,
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
func InitRoleRepo(db *gorm.DB) repositories.RoleRepo {
	return repositories.NewRoleRepo(db)
}
func InitPermissionRepo(db *gorm.DB) repositories.PermissionRepo {
	return repositories.NewPermissionRepo(db)
}
