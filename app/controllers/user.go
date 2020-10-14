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

func UserCreate(r repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}
		if err := r.Create(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return

		}
		c.JSON(http.StatusCreated, gin.H{
			"message": user,
		})
	}
}
