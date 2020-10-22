package services

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// USER_KEY : The session key to grab current logged in user
const USER_KEY string = "USER_KEY"

// SetSession : Adds a session key and value to the session
func SetSession(key string, value interface{}, c *gin.Context) {
	session := sessions.Default(c)

	session.Set(key, value)
	if err := session.Save(); err != nil {
		RespondInternalServerError(c, err, "Failed to save session")
		return
	}
}

// SessionRemove : Removes a session key from the session
func SessionRemove(key string, c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(key)
}
