package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"

	actionsDatabase "github.com/electronlabs/vibes-api/actions/database"
	actionsRoutes "github.com/electronlabs/vibes-api/actions/routes"
	actionsService "github.com/electronlabs/vibes-api/actions/service"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(mongo *mongo.Client) http.Handler {
	router := gin.Default()
	HealthRoutes(router)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	actionsGroup := v1.Group("/actions")
	actionsRepo := actionsDatabase.New(mongo)
	actionsSvc := actionsService.New(actionsRepo)
	actionsRoutes.CreateHandlers(actionsGroup, actionsSvc)

	return router
}
