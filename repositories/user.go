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
	Find(params UserFindParams) (models.User, error)
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

type UserFindParams struct {
	ID       uint
	Email    string
	Username string
}

func (u userRepo) Find(p UserFindParams) (models.User, error) {
	var user models.User
	if p.ID != 0 {
		tx := u.DB.First(&user, p.ID)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	} else if p.Email != "" {
		tx := u.DB.Where("email = ?", p.Email).First(&user)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	} else if p.Username != "" {
		tx := u.DB.Where("username = ?", p.Username).First(&user)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	}

	return user, nil
}
