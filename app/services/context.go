package services

import (
	"discord-clone-server/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserContext(c *gin.Context) (models.User, error) {
	userContext, ok := c.Get(USER_KEY)
	if !ok {
		return models.User{}, errors.New("no user on context")
	}

	user, ok := userContext.(models.User)
	if !ok {
		return models.User{}, errors.New("Cant convert user from context")
	}
	return user, nil
}
