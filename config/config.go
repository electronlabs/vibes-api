package config

import (
	"github.com/electronlabs/vibes-api/utils/env"
)

// Config is a struct that contains configuration variables
type Config struct {
	Environment         string
	Port                string
	MongoURI            string
	JWKSURL             string
	Auth0NativeClientID string
	Auth0NativeSecret   string
	Auth0NativeIssuer   string
	Auth0APIAudience    string
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	env.CheckDotEnv()
	return &Config{
		Environment:         env.MustGet("ENV"),
		Port:                env.MustGet("PORT"),
		MongoURI:            env.MustGet("MONGO_URI"),
		JWKSURL:             env.MustGet("JWKS_URL"),
		Auth0NativeClientID: env.MustGet("AUTH0_NATIVE_CLIENT_ID"),
		Auth0NativeSecret:   env.MustGet("AUTH0_NATIVE_SECRET"),
		Auth0NativeIssuer:   env.MustGet("AUTH0_NATIVE_ISSUER"),
		Auth0APIAudience:    env.MustGet("AUTH0_API_AUDIENCE"),
	}
}
