package models

import (
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Name    string `json:"name"`
	Permissions   []*Permissions `gorm:"many2many:role_permissions"`
}
