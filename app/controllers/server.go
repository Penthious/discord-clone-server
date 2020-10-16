package controllers

import (
	"discord-clone-server/models"
	"discord-clone-server/repositories"

	"github.com/gin-gonic/gin"
)

type serverCreateParams struct {
	Name    string `json:"name" binding:"required,min=3,max=25"`
	Private bool   `json:"private"`
}

func ServerCreate(r repositories.ServerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUserContext(c)
		if err != nil {
			RespondBadRequestError(c, err, "Error getting user from context")
			return
		}
		var p serverCreateParams
		if err := c.ShouldBind(&p); err != nil {
			RespondBadRequestError(c, err, err.Error())
			return
		}

		server := models.Server{
			Name:    p.Name,
			Private: p.Private,
		}
		if err := r.Create(&server); err != nil {
			RespondBadRequestError(c, err, "Error creating server")
			return
		}

		if err := r.Append(&user, server); err != nil {
			RespondBadRequestError(c, err, "Error setting user to server")
			return
		}

		RespondStatusCreated(c, "server", server)
		return
	}
}
