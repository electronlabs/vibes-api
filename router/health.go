package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthRoutes create and returns the health http routes
func HealthRoutes(router *gin.Engine) {

	health := router.Group("/health")
	{
		health.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "health ok")
		})
	}
}
