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
	GetUserServerRoles(userID uint, serverID uint) ([]models.Role, error)
	AttachServerRoles([]models.ServerUserRole) (models.User, error)
	CreateAdminRole() (models.Role, error)
	CreateBaseRole() (models.Role, error)
	GetServerRoles(server *models.Server) ([]models.Role, error)
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

// GetUserServerRoles : Finds all roles of a user for a given server
func (r roleRepo) GetUserServerRoles(userID uint, serverID uint) ([]models.Role, error) {
	query := `
	SELECT r.name, r.id, r.server_id FROM servers AS s
	JOIN server_users AS su on s.id = su.server_id
	join server_user_roles sur on su.user_id = sur.user_id
	join roles as r on sur.role_id = r.id
	where su.server_id = ?
	and sur.server_id = ?
	and sur.user_id = ?
`

	var userRoles []models.Role
	r.DB.Raw(query, serverID, serverID, userID).Scan(&userRoles)

	return userRoles, nil
}

// AttachServerRoles : Attaches roles to a server
func (r roleRepo) AttachServerRoles(sur []models.ServerUserRole) (models.User, error) {

	var user models.User
	for _, s := range sur {
		r.DB.Exec("INSERT INTO `server_user_roles` (`server_id`, `user_id`, `role_id`) VALUES (?, ?, ?)", s.ServerID, s.UserID, s.RoleID)

	}
	return user, nil
}

func (r roleRepo) GetServerRoles(server *models.Server) ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Model(&server).Association("Roles").Find(&roles)
	return roles, err
}

func (r roleRepo) CreateAdminRole() (models.Role, error) {
	var permissions []*models.Permission
	permissionsList := []string{"admin"}
	tx := r.DB.Where("permission = ?", permissionsList).First(&permissions)
	if err := tx.Error; err != nil {
		return models.Role{}, tx.Error
	}

	role := models.Role{
		Name:        "Admin",
		Permissions: permissions,
	}

	return role, tx.Error
}

func (r roleRepo) CreateBaseRole() (models.Role, error) {
	var permissions []*models.Permission
	permissionsList := []string{
		"create_invite",
		"manage_emojis",
		"read_channels",
		"send_message",
		"embed_link",
		"attach_file",
		"read_message_history",
		"external_emojis",
		"use_mentions",
		"add_reaction",
		"connect",
		"speak",
		"video",
	}
	tx := r.DB.Where("permission IN ?", permissionsList).Find(&permissions)
	if err := tx.Error; err != nil {
		return models.Role{}, tx.Error
	}

	role := models.Role{
		Name:        "Base",
		Permissions: permissions,
	}

	return role, tx.Error
}
