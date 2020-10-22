package app

import (
	"discord-clone-server/app/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// InitRouter : Sets up all routes
func InitRouter(s Services) (*gin.Engine, error) {

	r := gin.Default()
	r.Use(sessions.Sessions("discord_clone_session", sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	BaseRoutes(r, s)
	ServerRoutes(r, s)
	UserRoutes(r, s)

	return r, nil
}

// BaseRoutes : Sets up routes that dont really belong to anything
func BaseRoutes(r *gin.Engine, s Services) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controllers.Login(s.UserRepo))
	r.POST("/logout", controllers.Logout)

}

// ServerRoutes : Sets up the server routes
func ServerRoutes(r *gin.Engine, s Services) {
	server := r.Group("/servers")
	server.Use(AuthMiddleware(s.UserRepo))
	{
		server.POST("/", controllers.ServerCreate(s.ServerRepo, s.PermissionRepo, s.UserRepo, s.RoleRepo))
		server.POST("/invite", controllers.InviteUser(s.ServerRepo, s.PermissionRepo, s.Redis))
		server.POST("/join", controllers.JoinServer(s.ServerRepo, s.RoleRepo, s.Redis))
	}
}

// UserRoutes : Sets up the user routes
func UserRoutes(r *gin.Engine, s Services) {
	r.GET("/users", controllers.UserIndex(s.UserRepo))
	r.POST("/users", controllers.UserCreate(s.UserRepo))
}
