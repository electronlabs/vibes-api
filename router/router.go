package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	actions "github.com/electronlabs/vibes-api/actions/routes"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler() http.Handler {
	router := gin.Default()
	HealthRoutes(router)
	actions.NewActionsRoutesFactory()
	return router
}
