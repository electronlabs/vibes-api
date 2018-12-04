package main

import (
	"github.com/electronlabs/vibes-api/router"
	"github.com/electronlabs/vibes-api/utils/env"
)

func main() {
	port := env.MustGet("PORT")
	router.Start(port)
}
