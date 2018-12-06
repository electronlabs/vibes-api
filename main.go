package main

import (
	"net/http"

	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/database/mongodb"
	"github.com/electronlabs/vibes-api/router"
)

func main() {
	config := config.NewConfig()
	_, err := mongodb.Connect(config.MongoURI)
	if err != nil {
		panic(err)
	}

	router := router.NewHTTPHandler()

	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		panic(err)
	}
}
