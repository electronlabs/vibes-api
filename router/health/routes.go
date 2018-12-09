package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthRoutes create and returns the health http routes
func Routes(router *gin.Engine) {

	health := router.Group("/health")
	{
		health.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "health ok")
		})
	}
}
