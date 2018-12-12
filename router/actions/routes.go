package actions

import (
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/gin-gonic/gin"
)

// NewRoutesFactory create and returns a factory to create routes for the actions
func NewRoutesFactory(service *actions.Service) func(group *gin.RouterGroup) {
	actionRoutesFactory := func(group *gin.RouterGroup) {

		group.GET("/", func(c *gin.Context) {
			actionList, _ := service.ListActions()
			c.JSON(http.StatusOK, actionList)
		})
	}

	return actionRoutesFactory
}
