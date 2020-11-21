package models

import (
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Name       string `json:"name"`
	Private    bool   `json:"private"`
	User_ID    uint
	Users      []*User `gorm:"many2many:server_users"`
	Roles      []Role
	Categories []Category
}
