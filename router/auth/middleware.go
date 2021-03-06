package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenValidator struct
type TokenValidator interface {
	CheckJWT(tokenStr string) (*User, error)
}

// jwtFromAuthHeader takes a request and extracts the JWT token from the Authorization header.
func jwtFromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// New creates an auth middleware
func New(validator TokenValidator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := jwtFromAuthHeader(ctx.Request)
		if err != nil {
			ctx.AbortWithStatus(400)
		}

		user, err := validator.CheckJWT(tokenStr)
		if err != nil {
			ctx.AbortWithStatus(401)
		}

		ctx.Set("user", user)
	}
}
