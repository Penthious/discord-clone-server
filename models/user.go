package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Servers   []*Server `gorm:"many2many:server_users"`
	Roles     []*Role   `gorm:"many2many:server_user_roles"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := user.CheckPassword("password0"); err != nil { password error }
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func HashPassword(password string) (string, error) {
	fmt.Println(password)

	if password != "" {
		hash, err := makePassword(password)
		if err != nil {
			return "", err
		}
		password = hash
		return hash, nil
	}

	return "", nil
}

func makePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
