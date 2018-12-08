package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ActionsRoutes create and returns the actions http routes
func ActionsRoutes(router *gin.Engine) {

	actions := router.Group("/actions")
	{
		actions.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "actions")
		})
	}
}
