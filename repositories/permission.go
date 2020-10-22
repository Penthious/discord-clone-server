package repositories

import (
	"discord-clone-server/models"
	"errors"
	"fmt"

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
	GetUserServerPermissions(uint, uint) ([]models.Permission, error)
	CanAccess(required []models.Permission, userPerms []models.Permission) error
	InviteUserPermission() ([]models.Permission, error)
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

// GetUserServerPermissions : Finds all permissions a user has for a given server
func (r permissionRepo) GetUserServerPermissions(userID uint, serverID uint) ([]models.Permission, error) {
	query := `
SELECT p.name, p.detail, p.permission, p.id FROM servers AS s
JOIN server_users AS su on s.id = su.server_id
join server_user_roles sur on s.id = sur.server_id
join roles as r on sur.role_id = r.id
join role_permissions rp on r.id = rp.role_id
join permissions p on rp.permission_id = p.id
where su.user_id = 1
and r.server_id = 1
group by p.id
`

	var userPermissions []models.Permission
	r.DB.Raw(query, userID, serverID).Scan(&userPermissions)

	fmt.Printf("userRoles: %v\n", userPermissions)

	return userPermissions, nil
}

func (r permissionRepo) CanAccess(requiredPermissions []models.Permission, userPermissions []models.Permission) error {
	for _, p := range userPermissions {
		if p.Permission == "admin" {
			return nil
		}
	}

	var found []bool
	for _, p := range userPermissions {
		for _, r := range requiredPermissions {
			if p.Permission == r.Permission {
				found = append(found, true)
			}
		}
	}
	if len(found) != len(requiredPermissions) {
		return errors.New("Permission's mismatch")
	}

	return nil
}

func (r permissionRepo) InviteUserPermission() ([]models.Permission, error) {
	var perms []models.Permission
	var canInvitePerms = []string{"can_invite"}
	err := r.DB.Where("permission IN ?", canInvitePerms).Find(&perms)

	return perms, err.Error
}
