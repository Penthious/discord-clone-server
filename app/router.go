package app

import (
	"fmt"
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
	r.POST("/users", func(c *gin.Context) {
		var request struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "something went wrong binding request"})
			// RespondBadRequestError(c, err, "error binding set request store", s.log)
			return
		}

		fmt.Println("-------")
		fmt.Printf("user: %#v\n", request)
		fmt.Println("-------")

		c.JSON(http.StatusCreated, gin.H{
			"message": "user created",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		users, err := s.UserRepo.Get()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	return r, nil
}
