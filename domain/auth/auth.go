package auth

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

// validateClaims validates a token's time based claims, audience and issuer.
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

// getPublicKey searches inside a JWKS for a key corresponding to the token passed
// and, if it's found, it generates the corresponding public RSA key using the key info.
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

// tokenVerifier returns the function that verifies the claims and generates the
// public RSA key to validate the passed JWT against.
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

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions.
func CheckJWT(tokenStr string, jwksURL string, audience string, issuer string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, tokenVerifier(jwksURL, audience, issuer))
	if err != nil {
		return nil, err
	}

	return token, nil
}
