package auth

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
)

// jwtFromAuthHeader takes a request and extracts the JWT token from the Authorization header.
func jwtFromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func tokenVerifier(jwksURL string) func(token *jwt.Token) (interface{}, error) {
	set, err := jwk.FetchHTTP(jwksURL)

	return func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			return nil, err
		}

		keyID, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("expecting JWT header to have string kid")
		}

		if key := set.LookupKeyID(keyID); len(key) == 1 {
			return key[0].Materialize()
		}

		return nil, errors.New("Unable to find key")
	}
}

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions for the request.
func CheckJWT(jwksURL string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")

		tokenStr := ""
		if substr := strings.Split(header, " "); len(substr) == 2 {
			tokenStr = substr[1]
		}

		tokenStr, err := jwtFromAuthHeader(ctx.Request)
		if err != nil {
			ctx.AbortWithStatus(400)
		}

		token, err := jwt.Parse(tokenStr, tokenVerifier(jwksURL))
		if err != nil {
			ctx.AbortWithStatus(401)
		}

		ctx.Set("tokenClaims", token.Claims)
		ctx.Next()
	}
}
