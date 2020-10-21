package app

import (
	"discord-clone-server/app/services"
	"discord-clone-server/repositories"
	"errors"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : Middleware to check if user is logged in via session
func AuthMiddleware(ur repositories.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userCookieID := session.Get(services.USER_KEY)
		if userCookieID == nil {
			services.RespondUnauthorizedError(c, errors.New("User cookie not found"), "unauthorized")
			return
		}
		if userCookieID == 0 {
			services.RespondUnauthorizedError(c, errors.New("User cookie is 0"), "unauthorized")
			return
		}
		userFindParams := repositories.UserFindParams{ID: userCookieID.(uint)}
		user, err := ur.Find(userFindParams)
		if err != nil {
			services.RespondUnauthorizedError(c, err, "unauthorized")
			return
		}
		c.Set(services.USER_KEY, user)
		c.Next()
	}
}
