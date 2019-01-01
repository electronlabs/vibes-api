package main

import (
	"net/http"

	"github.com/electronlabs/vibes-api/utils/token"

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

	tokenConfig := &token.Config{Audience: configuration.Auth.Audience, Issuer: configuration.Auth.Issuer, JwksUrl: configuration.Auth.JWKSURL}
	validator := token.NewValidator(tokenConfig)

	actionsRepo := actionsStore.New(mongo)
	actionsSvc := actions.NewService(actionsRepo)

	httpRouter := router.NewHTTPHandler(actionsSvc, validator)

	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
