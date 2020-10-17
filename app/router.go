package app

import (
	"discord-clone-server/app/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// thing
func InitRouter(s Services) (*gin.Engine, error) {

	r := gin.Default()
	r.Use(sessions.Sessions("discord_clone_session", sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controllers.Login(s.UserRepo))
	r.POST("/logout", controllers.Logout)

	r.GET("/users", controllers.UserIndex(s.UserRepo))
	r.POST("/users", controllers.UserCreate(s.UserRepo))

	auth := r.Group("/auth")
	auth.Use(AuthMiddleware(s.UserRepo))
	{
		auth.GET("/status", status)
		auth.POST("/servers", controllers.ServerCreate(s.ServerRepo, s.PermissionRepo))
	}
	return r, nil
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
