package app

import (
	"discord-clone-server/app/controllers"
	"discord-clone-server/repositories"
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
	r.GET("/users", controllers.UserIndex(s.UserRepo))
	r.POST("/users", controllers.UserCreate(s.UserRepo))

	auth := r.Group("/auth")
	auth.Use(AuthRequired(s.UserRepo))
	{
		auth.GET("/status", status)
	}

	return r, nil
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func AuthRequired(ur repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userCookieID := session.Get("user").(uint)
		if userCookieID == 0 {
			// Abort the request with the appropriate error code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		user, err := ur.Find(userCookieID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("user", user)
		// Continue down the chain to handler etc
		c.Next()
	}
}
