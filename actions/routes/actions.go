package routes

import (
	"github.com/electronlabs/vibes-api/actions/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewActionsRoutesFactory create and returns a factory to create routes for the actions
func NewActionsRoutesFactory(service *service.ActionsService) func(*gin.Engine) {
	actionRoutesFactory := func(router *gin.Engine) {

		ac := router.Group("/actions")
		{
			ac.GET("/", func(c *gin.Context) {
				actionList, _ := service.GetActions()
				c.JSON(http.StatusOK, actionList)
			})
		}
	}

	return actionRoutesFactory
}
