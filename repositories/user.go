package repositories

import (
	"discord-clone-server/models"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		DB: db,
	}
}

type UserRepo interface {
	Create(models.User) error
}

type userRepo struct {
	DB *gorm.DB
}

func (u userRepo) Create(user models.User) error {
	return u.DB.Create(&user).Error
}
