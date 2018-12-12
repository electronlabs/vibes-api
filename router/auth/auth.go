package auth

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
)

const (
	jwksURL = ""
)

func verifyToken(token *jwt.Token) (interface{}, error) {
	set, err := jwk.FetchHTTP(jwksURL)
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

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions for the request.
func CheckJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")

		tokenStr := ""
		if substr := strings.Split(header, " "); len(substr) == 2 {
			tokenStr = substr[1]
		}

		token, err := jwt.Parse(tokenStr, verifyToken)
		if err != nil {
			// Return 401
		}

		ctx.Set("tokenClaims", token.Claims)

		ctx.Next()
	}
}
