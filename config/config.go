package config

import (
	"github.com/electronlabs/vibes-api/utils/env"
)

// Auth contains variables required for authenticaton
type Auth struct {
	JWKSURL  string
	Audience string
	Issuer   string
}

// Config is a struct that contains configuration variables
type Config struct {
	Environment string
	Port        string
	MongoURI    string
	Auth        *Auth
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	env.CheckDotEnv()
	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        env.MustGet("PORT"),
		MongoURI:    env.MustGet("MONGO_URI"),
		Auth: &Auth{
			JWKSURL:  env.MustGet("JWKS_URL"),
			Audience: env.MustGet("JWT_AUDIENCE"),
			Issuer:   env.MustGet("JWT_ISSUER"),
		},
	}
}
