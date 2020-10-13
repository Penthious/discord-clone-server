package app

import (
	"discord-clone-server/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// thing
func InitRouter(s Services) (*gin.Engine, error) {

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", controllers.UserIndex(s.UserRepo))
	r.POST("/users", controllers.UserCreate(s.UserRepo))

	return r, nil
}
