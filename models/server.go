package models

import (
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Name    string  `json:"name"`
	Private bool    `json:"private"`
	Users   []*User `gorm:"many2many:server_users"`
}
