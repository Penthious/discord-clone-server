package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const USER_KEY = "USER_KEY"

func SetSession(key string, value interface{}, c *gin.Context) {
	session := sessions.Default(c)

	session.Set(key, value)
	if err := session.Save(); err != nil {
		RespondInternalServerError(c, err, "Failed to save session")
		return
	}
}

func SessionRemove(key string, c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(key)
}
