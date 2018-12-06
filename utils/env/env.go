package env

import "os"

// MustGet will return the env or panic if not present.
func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("Env key missing " + key)
	}
	return val
}
