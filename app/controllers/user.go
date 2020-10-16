package controllers

import (
	"discord-clone-server/models"
	"discord-clone-server/repositories"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
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
		session := sessions.Default(c)

		var p userCreateParams
		if err := c.ShouldBind(&p); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}

		session.Set("user", user.ID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": user,
		})
	}
}

type loginParams struct {
	Username string `json:"username" binding:"required_without=Email"`
	Email    string `json:"email" binding:"required_without=Username"`
	Password string `json:"password" binding:"required"`
}

func Login(ur repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
