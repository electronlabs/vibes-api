package main

import (
	"github.com/electronlabs/vibes-api/utils/jwks"
	"net/http"

	actionsStore "github.com/electronlabs/vibes-api/data/actions"
	"github.com/electronlabs/vibes-api/data/shared/mongodb"
	"github.com/electronlabs/vibes-api/domain/actions"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/router"
	"github.com/electronlabs/vibes-api/utils/validator"
)

func main() {
	configuration := config.NewConfig()

	mongo, err := mongodb.Connect(configuration.MongoURI)
	if err != nil {
		panic(err)
	}

	jwksToken, err := jwks.New(configuration.Auth.JWKSURL)
	if err != nil {
		panic(err)
	}

	validatorConfig := &validator.Config{Audience: configuration.Auth.Audience, Issuer: configuration.Auth.Issuer}
	tokenValidator := validator.New(validatorConfig, jwksToken)

	actionsRepo := actionsStore.New(mongo)
	actionsSvc := actions.NewService(actionsRepo)

	httpRouter := router.NewHTTPHandler(actionsSvc, tokenValidator)

	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}
}
