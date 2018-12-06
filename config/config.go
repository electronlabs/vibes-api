package config

import "github.com/electronlabs/vibes-api/utils/env"

// Config is a struct that contains configuration variables
type Config struct {
	Port string
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	port := env.MustGet("PORT")

	return &Config{
		Port: port,
	}
}
