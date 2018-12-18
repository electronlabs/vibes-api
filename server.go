package main

import (
	"net/http"

	actionsStore "github.com/electronlabs/vibes-api/data/actions"
	jwks "github.com/electronlabs/vibes-api/data/auth"
	"github.com/electronlabs/vibes-api/data/shared/mongodb"
	"github.com/electronlabs/vibes-api/domain/actions"
	"github.com/electronlabs/vibes-api/domain/auth"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/router"
)

func main() {
	configuration := config.NewConfig()

	mongo, err := mongodb.Connect(configuration.MongoURI)
	if err != nil {
		panic(err)
	}

	authRepo, err := jwks.New(configuration.Auth.JWKSURL)
	if err != nil {
		panic(err)
	}
	authSvc := auth.NewService(authRepo, &auth.Config{Audience: configuration.Auth.Audience, Issuer: configuration.Auth.Issuer})

	actionsRepo := actionsStore.New(mongo)
	actionsSvc := actions.NewService(actionsRepo)

	httpRouter := router.NewHTTPHandler(authSvc, actionsSvc)

	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
