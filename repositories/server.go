package repositories

import (
	"discord-clone-server/models"

	"gorm.io/gorm"
)

func NewServerRepo(db *gorm.DB) ServerRepo {
	return serverRepo{
		DB: db,
	}
}

type ServerRepo interface {
	Create(*models.User, *models.Server) error
}

type serverRepo struct {
	DB *gorm.DB
}

func (s serverRepo) Create(user *models.User, server *models.Server) error {

	return s.DB.Model(&user).Association("Servers").Append(&server)
	// return s.DB.Model(&user).Association("Servers").DB.Create(&server).Error
	// return s.DB.Create(&server).Error
}
