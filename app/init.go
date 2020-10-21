package app

import (
	"context"
	"discord-clone-server/repositories"
	"discord-clone-server/utils"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ctx = context.Background()

type Services struct {
	DB             *gorm.DB
	Redis          *redis.Client
	UserRepo       repositories.UserRepo
	ServerRepo     repositories.ServerRepo
	RoleRepo       repositories.RoleRepo
	PermissionRepo repositories.PermissionRepo
}

/*
func initRedisClient() (*redis.Client, error) {
	redisDB, err := util.GetEnvIntE("REDIS_DB")
	if err != nil {
		return nil, errors.Wrap(err, "error getting configured redis DB")
	}
	cl := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_ADDR"),
		DB:          redisDB,
		DialTimeout: util.GetEnvDurationD("REDIS_DIAL_TIMEOUT", 100*time.Millisecond),
	})
	if err := cl.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "error connecting to redis")
	}
	return cl, nil
}
*/

func InitServices() (Services, error) {
	var s Services
	var err error

	s.DB, err = InitDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v \n", err.Error())
	}

	s.Redis, err = InitRedis()
	if err != nil {
		log.Fatalf("Error connecting to REDIS: %v \n", err.Error())
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

func InitRedis() (*redis.Client, error) {
	fmt.Printf("REDDIS ADDR: %v \n", os.Getenv("REDIS_ADDR"))
	cl := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
		// Addr:        os.Getenv("REDIS_ADDR"),
		// DB:          redisDB,
		// DialTimeout: util.GetEnvDurationD("REDIS_DIAL_TIMEOUT", 100*time.Millisecond),
	})
	if err := cl.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return cl, nil
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
