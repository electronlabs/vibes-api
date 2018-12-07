package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateHealthRoute create and returns the health http routes
func CreateHealthRoute(router *gin.Engine) {

	health := router.Group("/health")
	{
		health.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "health ok")
		})
	}
}
