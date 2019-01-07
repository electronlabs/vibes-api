package router

import (
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/electronlabs/vibes-api/router/auth"

	actionsRoutes "github.com/electronlabs/vibes-api/router/actions"
	healthRoutes "github.com/electronlabs/vibes-api/router/health"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(actionsSvc *actions.Service, validator auth.TokenValidator) http.Handler {
	router := gin.Default()

	authMid := auth.New(validator)

	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	api := router.Group("/api")

	actionsGroup := api.Group("/actions")
	actionsGroup.Use(authMid)
	actionsRoutes.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
