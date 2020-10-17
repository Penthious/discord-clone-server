package models

import (
	"gorm.io/gorm"
)

type Permissions struct {
	gorm.Model
	Name       string   `json:"name"`
	Permission string   `json:"permission"`
	Detail     string   `json:"detail"`
	Roles      []*Roles `gorm:"many2many:role_permissions"`
}
