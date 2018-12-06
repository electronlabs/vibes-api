package main

import (
	"github.com/electronlabs/vibes-api/config"
	"github.com/electronlabs/vibes-api/router"
)

func main() {
	config := config.NewConfig()
	router.Start(config.Port)
}
