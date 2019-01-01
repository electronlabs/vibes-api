package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TokenValidator interface {
	CheckJWT(tokenStr string) (*jwt.Token, error)
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

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions for the request.
func NewAuthMiddleware(validator TokenValidator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := jwtFromAuthHeader(ctx.Request)
		if err != nil {
			ctx.AbortWithStatus(400)
		}

		jwtToken, err := validator.CheckJWT(tokenStr)
		if err != nil {
			ctx.AbortWithStatus(401)
		}

		ctx.Set("user", jwtToken)
	}
}
