package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name       string  `json:"name"`
	Permission string  `json:"permission"`
	Detail     string  `json:"detail"`
	Type       string  `json:"type"`
	Roles      []*Role `gorm:"many2many:role_permissions"`
}
