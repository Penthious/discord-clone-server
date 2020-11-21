package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name"`
	ServerID    uint
	Permissions []*Permission `gorm:"many2many:role_permissions"`
	Users       []*User       `gorm:"many2many:server_user_roles"`
}
