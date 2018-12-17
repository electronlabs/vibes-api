package main

import (
	"net/http"

	actionsStore "github.com/electronlabs/vibes-api/data/actions"
	"github.com/electronlabs/vibes-api/data/shared/mongodb"
	"github.com/electronlabs/vibes-api/domain/actions"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/router"
)

func main() {
	configuration := config.NewConfig()

	mongo, err := mongodb.Connect(configuration.MongoURI)
	if err != nil {
		panic(err)
	}

	actionsRepo := actionsStore.New(mongo, configuration.DatabaseName)
	actionsSvc := actions.NewService(actionsRepo)

	httpRouter := router.NewHTTPHandler(actionsSvc)

	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
