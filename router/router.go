package router

import (
	"github.com/electronlabs/vibes-api/domain/actions"
	actionsRoute "github.com/electronlabs/vibes-api/router/actions"
	"github.com/electronlabs/vibes-api/router/health"
	"net/http"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/domain/actions"
	actionsRoutes "github.com/electronlabs/vibes-api/router/actions"
	healthRoutes "github.com/electronlabs/vibes-api/router/actions"
	"github.com/electronlabs/vibes-api/router/auth"

	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(authConfig *config.Auth, actionsSvc *actions.Service) http.Handler {
	router := gin.Default()
	authMiddleware := auth.CheckJWT(authConfig.JWKSURL, authConfig.Audience, authConfig.Issuer)

	healthGroup := router.Group("/health")
	healthGroup.Use(authMiddleware)
	healthRoutes.NewRoutesFactory(actionsSvc)(healthGroup)

	api := router.Group("/api")

	actionsGroup := api.Group("/actions")
	actionsGroup.Use(authMiddleware)
	actionsRoutes.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
