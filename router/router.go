package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	actions "github.com/electronlabs/vibes-api/actions/routes"
	users "github.com/electronlabs/vibes-api/users/routes"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler() http.Handler {
	router := gin.Default()
	HealthRoutes(router)
	actions.ActionsRoutes(router)
	users.UsersRoutes(router)
	return router
}
