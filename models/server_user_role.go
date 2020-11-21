package models

import (
	"gorm.io/gorm"
)

type ServerUserRole struct {
	gorm.Model
	ServerID uint
	UserID   uint
	RoleID   uint
}
