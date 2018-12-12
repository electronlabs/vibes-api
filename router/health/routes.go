package health

import (
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes to check API health
func NewRoutesFactory(service *actions.Service) func(group *gin.RouterGroup) {
	actionRoutesFactory := func(group *gin.RouterGroup) {
		group.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "health ok")
		})
	}

	return actionRoutesFactory
}
