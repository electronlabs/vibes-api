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

func validateClaims(token *jwt.Token, audience string, issuer string) error {
	// Validate time based claims
	err := token.Claims.Valid()
	if err != nil {
		return err
	}

	// Validate audience
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
	if !checkAud {
		return errors.New("invalid token audience")
	}

	// Validate issuer
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
	if !checkIss {
		return errors.New("invalid token issuer")
	}

	return nil
}

func getPublicKey(set *jwk.Set, token *jwt.Token) (interface{}, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if keys := set.LookupKeyID(keyID); len(keys) == 1 {
		return keys[0].Materialize()
	}

	return nil, errors.New("Unable to find key")
}

func tokenVerifier(jwksURL string, audience string, issuer string) func(token *jwt.Token) (interface{}, error) {
	set, err := jwk.FetchHTTP(jwksURL)

	return func(token *jwt.Token) (interface{}, error) {
		// Error fetching JWKS
		if err != nil {
			return nil, err
		}

		err = validateClaims(token, audience, issuer)
		if err != nil {
			return nil, err
		}

		return getPublicKey(set, token)
	}
}

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions for the request.
func CheckJWT(jwksURL string, audience string, issuer string) gin.HandlerFunc {
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

		token, err := jwt.Parse(tokenStr, tokenVerifier(jwksURL, audience, issuer))
		if err != nil {
			ctx.AbortWithStatus(401)
		}

		ctx.Set("tokenClaims", token.Claims)
		ctx.Next()
	}
}
