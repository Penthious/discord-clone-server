package controllers

import (
	"discord-clone-server/models"
	"discord-clone-server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverCreateParams struct {
	Name string `json:"name" binding:"required,min=3,max=25"`
}

func ServerCreate(r repositories.ServerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUserContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}
		var p serverCreateParams
		if err := c.ShouldBind(&p); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}

		server := models.Server{
			Name: p.Name,
		}
		if err := r.Create(&user, &server); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"server": server,
		})
	}
}
