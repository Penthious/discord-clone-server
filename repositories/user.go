package repositories

import (
	"discord-clone-server/models"
	"fmt"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		DB: db,
	}
}

type UserRepo interface {
	Create(models.User) error
	Get() ([]models.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func (u userRepo) Create(user models.User) error {
	fmt.Printf("user: %v\n", user)
	return u.DB.Create(&user).Error
}

func (u userRepo) Get() ([]models.User, error) {
	var users []models.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
