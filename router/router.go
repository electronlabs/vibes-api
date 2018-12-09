package router

import (
	"github.com/electronlabs/vibes-api/domain/actions"
	actionsRoute "github.com/electronlabs/vibes-api/router/actions"
	"github.com/electronlabs/vibes-api/router/health"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(actionsSvc *actions.Service) http.Handler {
	router := gin.Default()
	health.Routes(router)

	api := router.Group("/api")

	actionsGroup := api.Group("/actions")
	actionsRoute.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
