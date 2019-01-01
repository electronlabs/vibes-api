package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	invalidAudience = "invalid token audience"
	invalidIssuer   = "invalid token issuer"
	missingKID      = "expecting JWT header to have string kid"
)

// Config struct defines auth service configuration variables
type Config struct {
	Audience string
	Issuer   string
	JwksUrl  string
}

type Validator struct {
	audience string
	issuer   string
	jwksUrl  string
}

// NewValidator creates a new instance of token validator
func NewValidator(config *Config) *Validator {
	return &Validator{
		audience: config.Audience,
		issuer:   config.Issuer,
		jwksUrl:  config.JwksUrl,
	}
}

// validateClaims validates a token's time based claims, audience and issuer.
func (validator *Validator) validateClaims(token *jwt.Token) error {
	// Validate time based claims
	err := token.Claims.Valid()
	if err != nil {
		return err
	}

	// Validate audience
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(validator.audience, false)
	if !checkAud {
		return errors.New(invalidAudience)
	}

	// Validate issuer
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(validator.issuer, false)
	if !checkIss {
		return errors.New(invalidIssuer)
	}

	return nil
}

// getPublicKeyID extracts the public key ID from the passed token.
func getPublicKeyID(token *jwt.Token) (string, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return "", errors.New(missingKID)
	}

	return keyID, nil
}

// tokenVerifier returns the function that verifies the claims and generates the
// public RSA key to validate the passed JWT against.
func (validator *Validator) tokenVerifier() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		err := validator.validateClaims(token)
		if err != nil {
			return nil, err
		}

		keyID, err := getPublicKeyID(token)
		if err != nil {
			return nil, err
		}

		jwks, err := NewJWKS(validator.jwksUrl)
		if err != nil {
			return nil, err
		}

		return jwks.getPublicKey(keyID)
	}
}

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions.
func (validator *Validator) CheckJWT(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, validator.tokenVerifier())
	if err != nil {
		return nil, err
	}

	return token, nil
}
