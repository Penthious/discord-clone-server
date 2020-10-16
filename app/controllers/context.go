package controllers

import (
	"discord-clone-server/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func getUserContext(c *gin.Context) (models.User, error) {
	userContext, ok := c.Get("user")
	if !ok {
		return models.User{}, errors.New("no user on context")
	}

	user, ok := userContext.(models.User)
	if !ok {
		return models.User{}, errors.New("no user on context")
	}
	return user, nil
}
