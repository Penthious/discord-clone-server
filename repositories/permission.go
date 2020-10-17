package repositories

import (
	"discord-clone-server/models"

	"gorm.io/gorm"
)

func NewPermissionRepo(db *gorm.DB) PermissionRepo {
	return permissionRepo{
		DB: db,
	}
}

type PermissionRepo interface {
	Create(*models.Permission) error
	Get() ([]models.Permission, error)
	Find(params PermissionFindParams) (models.Permission, error)
}

type permissionRepo struct {
	DB *gorm.DB
}

func (r permissionRepo) Create(permission *models.Permission) error {
	return r.DB.Create(&permission).Error
}

func (r permissionRepo) Get() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.DB.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

type PermissionFindParams struct {
	ID         uint
	Permission string
}

func (r permissionRepo) Find(p PermissionFindParams) (models.Permission, error) {
	var permission models.Permission
	if p.ID != 0 {
		tx := r.DB.First(&permission, p.ID)
		if err := tx.Error; err != nil {
			return models.Permission{}, tx.Error
		}
	} else if p.Permission != "" {
		tx := r.DB.Where("permission = ?", p.Permission).First(&permission)
		if err := tx.Error; err != nil {
			return models.Permission{}, tx.Error
		}
	}

	return permission, nil
}
