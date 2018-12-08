package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersRoutes create and returns the users http routes
func UsersRoutes(router *gin.Engine) {

	users := router.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "users")
		})
	}
}
