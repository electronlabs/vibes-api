package routes

import (
	"net/http"

	"github.com/electronlabs/vibes-api/actions/service"

	"github.com/gin-gonic/gin"
)

// CreateHandlers registers action route handlers
func CreateHandlers(group *gin.RouterGroup, service *service.Service) {
	group.GET("/", func(ctx *gin.Context) {
		actionList, _ := service.GetActions()
		ctx.JSON(http.StatusOK, actionList)
	})
}
