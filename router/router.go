package router

import (
	"net/http"

	actions "github.com/electronlabs/vibes-api/actions/router"
	users "github.com/electronlabs/vibes-api/users/router"

	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler() http.Handler {
	router := gin.Default()
	CreateHealthRoute(router)
	actions.CreateRoutes(router)
	users.CreateRoutes(router)
	return router
}
