package router

import (
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	actionsRoutes "github.com/electronlabs/vibes-api/router/actions"
	healthRoutes "github.com/electronlabs/vibes-api/router/actions"
	auth "github.com/electronlabs/vibes-api/router/middleware/authentication"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(jwksURL, issuer, audience string, actionsSvc *actions.Service) http.Handler {
	router := gin.Default()
	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory(actionsSvc)(healthGroup)

	api := router.Group("/api")
	actionsGroup := api.Group("/actions")
	actionsGroup.Use(auth.CheckJWT(jwksURL, issuer, audience))
	actionsRoutes.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
