package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondBadRequestError : Respond with a bad request
func RespondBadRequestError(c *gin.Context, err error, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
	fmt.Printf("error: %v\n", err.Error())
}

// RespondInternalServerError : Respond with an internal server error
func RespondInternalServerError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
	fmt.Printf("error: %v\n", err.Error())
}

// RespondStatusCreated : Respond with status created
func RespondStatusCreated(c *gin.Context, key string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		key: data,
	})
}

// RespondStatusOK : Respond with status ok
func RespondStatusOK(c *gin.Context, key string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		key: data,
	})
}

// RespondStatusAccepted : Respond with status accepted
func RespondStatusAccepted(c *gin.Context, key string, data interface{}) {
	c.JSON(http.StatusAccepted, gin.H{
		key: data,
	})
}
