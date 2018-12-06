package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler() http.Handler {
	router := gin.Default()
	return router
}
