package config

import (
	"github.com/electronlabs/vibes-api/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Port         string
	MongoURI     string
	DatabaseName string
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	env.CheckDotEnv()
	return &Config{
		Port:         env.MustGet("PORT"),
		MongoURI:     env.MustGet("MONGO_URI"),
		DatabaseName: env.MustGet("DATABASE_NAME"),
	}
}
