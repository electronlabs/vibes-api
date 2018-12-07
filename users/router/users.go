package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRoutes create and returns the users http routes
func CreateRoutes(router *gin.Engine) {

	users := router.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "users")
		})
	}
}
