package config

import (
	"github.com/electronlabs/vibes-api/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Environment string
	Port        string
	MongoURI    string
	JWKSURL     string
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	env.CheckDotEnv()
	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        env.MustGet("PORT"),
		MongoURI:    env.MustGet("MONGO_URI"),
		JWKSURL:     env.MustGet("JWKS_URL"),
	}
}
