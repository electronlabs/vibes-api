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

		group.GET("/:actionId", func(c *gin.Context) {
			actionId := c.Param("actionId")
			action, err := service.GetAction(actionId)
			if err != nil {
				c.String(404, "Oops! Doc not found.")
				c.Error(err)
				c.AbortWithStatus(404)
			} else {
				c.JSON(http.StatusOK, action)
			}
		})
	}

	return actionRoutesFactory
}
