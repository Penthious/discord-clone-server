package repositories

import (
	"discord-clone-server/models"

	"gorm.io/gorm"
)

func NewRoleRepo(db *gorm.DB) RoleRepo {
	return roleRepo{
		DB: db,
	}
}

type RoleRepo interface {
	Create(*models.Role) error
	Get() ([]models.Role, error)
	Find(params RoleFindParams) (models.Role, error)
}

type roleRepo struct {
	DB *gorm.DB
}

func (r roleRepo) Create(role *models.Role) error {
	return r.DB.Create(&role).Error
}

func (r roleRepo) Get() ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

type RoleFindParams struct {
	ID   uint
	Role string
}

func (r roleRepo) Find(p RoleFindParams) (models.Role, error) {
	var role models.Role
	if p.ID != 0 {
		tx := r.DB.First(&role, p.ID)
		if err := tx.Error; err != nil {
			return models.Role{}, tx.Error
		}
	} else if p.Role != "" {
		tx := r.DB.Where("role = ?", p.Role).First(&role)
		if err := tx.Error; err != nil {
			return models.Role{}, tx.Error
		}
	}

	return role, nil
}
