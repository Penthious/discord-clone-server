package controllers

import (
	"discord-clone-server/models"
	"discord-clone-server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserIndex(r repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := r.Get()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

type userCreateParams struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required,min=4,max=15"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=36"`
}

func UserCreate(r repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p userCreateParams
		if err := c.ShouldBind(&p); err != nil {
			RespondBadRequestError(c, err, err.Error())
			return
		}

		user := models.User{
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Username:  p.Username,
			Email:     p.Email,
			Password:  p.Password,
		}
		if err := r.Create(&user); err != nil {
			RespondBadRequestError(c, err, "Error creating user")
			return
		}

		SetSession(USER_KEY, user.ID, c)
		RespondStatusCreated(c, "user", user)
	}
}

type loginParams struct {
	Username string `json:"username" binding:"required_without=Email"`
	Email    string `json:"email" binding:"required_without=Username"`
	Password string `json:"password" binding:"required"`
}

func Login(ur repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p loginParams
		if err := c.ShouldBind(&p); err != nil {
			RespondBadRequestError(c, err, err.Error())
			return
		}
		userFindParams := repositories.UserFindParams{Email: p.Email, Username: p.Username}
		user, err := ur.Find(userFindParams)
		if err != nil {
			RespondBadRequestError(c, err, "Error finding user")
			return

		}

		if err := user.CheckPassword(p.Password); err != nil {
			RespondBadRequestError(c, err, "Error password or email mismatch")
		}

		SetSession(USER_KEY, user.ID, c)

		RespondStatusOK(c, "user", user)
		return
	}
}

func Logout(c *gin.Context) {
	SessionRemove(USER_KEY, c)
	RespondStatusAccepted(c, "message", "ok")
	return
}
