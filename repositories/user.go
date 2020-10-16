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
	Create(*models.User) error
	Get() ([]models.User, error)
	Find(id uint) (models.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func (u userRepo) Create(user *models.User) error {
	return u.DB.Create(&user).Error
}

func (u userRepo) Get() ([]models.User, error) {
	var users []models.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepo) Find(id uint) (models.User, error) {
	var user models.User
	tx := u.DB.First(&user, id)

	if err := tx.Error; err != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}
