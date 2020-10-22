package repositories

import (
	"discord-clone-server/models"
	"errors"

	"gorm.io/gorm"
)

func NewServerRepo(db *gorm.DB) ServerRepo {
	return serverRepo{
		DB: db,
	}
}

type ServerRepo interface {
	Create(*models.Server) error
	Append(*models.User, models.Server) error
	Find(uint, *models.Server) error
	UserExistsOnServer(*models.Server, models.User) error
}

type serverRepo struct {
	DB *gorm.DB
}

func (r serverRepo) Append(user *models.User, server models.Server) error {
	return r.DB.Model(&user).Association("Servers").Append(&server)
}

func (r serverRepo) Create(server *models.Server) error {
	return r.DB.Create(&server).Error
}

func (r serverRepo) Find(serverID uint, server *models.Server) error {
	return r.DB.Where("id = ?", serverID).Find(&server).Error
}

func (r serverRepo) UserExistsOnServer(server *models.Server, user models.User) error {
	var u []models.User
	if err := r.DB.Model(&server).Association("Users").Find(&u, user.ID); err != nil {
		return err
	}

	if len(u) != 1 {
		return errors.New("User not associated to server")
	}

	return nil
}
