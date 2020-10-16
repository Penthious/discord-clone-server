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
	Servers   []*Server `gorm:"many2many:server_user"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	hashedPassword, err := hashPassword(u.Password)
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

func hashPassword(password string) (string, error) {
	fmt.Println("before create")
	fmt.Println(password)

	if password != "" {
		hash, err := MakePassword(password)
		if err != nil {
			return "", err
		}
		password = hash
		return hash, nil
	}

	return "", nil
}

func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
