package models

type Server struct {
	Name  string
	Users []*User `gorm:"many2many:server_user"`
}
