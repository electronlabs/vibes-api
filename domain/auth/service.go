package auth

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	invalidAudience = "invalid token audience"
	invalidIssuer   = "invalid token issuer"
	missingKID      = "expecting JWT header to have string kid"
)

// AuthService interface defines authentication service behavior
type AuthService interface {
	CheckJWT(tokenStr string) (*jwt.Token, error)
}

// Config struct defines auth service configuration variables
type Config struct {
	Audience string
	Issuer   string
}

// Service is the struct that implements AuthService interface
type Service struct {
	repository AuthRepository
	audience   string
	issuer     string
}

// NewService creates a new instance of authentication Service
func NewService(repository AuthRepository, config *Config) *Service {
	return &Service{
		repository: repository,
		audience:   config.Audience,
		issuer:     config.Issuer,
	}
}

// validateClaims validates a token's time based claims, audience and issuer.
func (svc *Service) validateClaims(token *jwt.Token) error {
	// Validate time based claims
	err := token.Claims.Valid()
	if err != nil {
		return err
	}

	// Validate audience
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(svc.audience, false)
	if !checkAud {
		return errors.New(invalidAudience)
	}

	// Validate issuer
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(svc.issuer, false)
	if !checkIss {
		return errors.New(invalidIssuer)
	}

	return nil
}

// getPublicKeyID extracts the public key ID from the passed token.
func (svc *Service) getPublicKeyID(token *jwt.Token) (string, error) {
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return "", errors.New(missingKID)
	}

	return keyID, nil
}

// tokenVerifier returns the function that verifies the claims and generates the
// public RSA key to validate the passed JWT against.
func (svc *Service) tokenVerifier() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		err := svc.validateClaims(token)
		if err != nil {
			return nil, err
		}

		keyID, err := svc.getPublicKeyID(token)
		if err != nil {
			return nil, err
		}

		return svc.repository.GetPublicKey(keyID)
	}
}

// CheckJWT checks the JSON Web Token and verifies it has the correct permissions.
func (svc *Service) CheckJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, svc.tokenVerifier())
	if err != nil {
		return nil, err
	}

	return token, nil
}
