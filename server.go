package main

import (
	"github.com/electronlabs/vibes-api/domain/actions"
	"net/http"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/router"
)

func main() {
	configuration := config.NewConfig()

	actionsRepo := actions.Store{}
	actionsSvc := actions.NewService(&actionsRepo)

	httpRouter := router.NewHTTPHandler(actionsSvc)

	err := http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
