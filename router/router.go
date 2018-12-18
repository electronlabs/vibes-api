package router

import (
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/electronlabs/vibes-api/domain/auth"
	actionsRoutes "github.com/electronlabs/vibes-api/router/actions"
	healthRoutes "github.com/electronlabs/vibes-api/router/actions"
	authMiddleware "github.com/electronlabs/vibes-api/router/middleware/auth"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(authSvc auth.AuthService, actionsSvc *actions.Service) http.Handler {
	router := gin.Default()
	authMid := authMiddleware.CheckJWT(authSvc)

	healthGroup := router.Group("/health")
	healthGroup.Use(authMid)
	healthRoutes.NewRoutesFactory(actionsSvc)(healthGroup)

	api := router.Group("/api")

	actionsGroup := api.Group("/actions")
	actionsGroup.Use(authMid)
	actionsRoutes.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
