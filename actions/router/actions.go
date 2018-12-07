package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRoutes create and returns the actions http routes
func CreateRoutes(router *gin.Engine) {

	actions := router.Group("/actions")
	{
		actions.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "actions")
		})
	}
}
