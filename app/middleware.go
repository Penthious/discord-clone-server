package app

import (
	"discord-clone-server/app/controllers"
	"discord-clone-server/repositories"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ur repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userCookieID := session.Get(controllers.USER_KEY)
		if userCookieID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if userCookieID == 0 {
			// Abort the request with the appropriate error code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		userFindParams := repositories.UserFindParams{ID: userCookieID.(uint)}
		user, err := ur.Find(userFindParams)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set(controllers.USER_KEY, user)
		// Continue down the chain to handler etc
		c.Next()
	}
}
