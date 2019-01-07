package jwks

import (
	"errors"

	"github.com/lestrrat-go/jwx/jwk"
)

// JWKS handles fetching JSON Web Key Set and the required keys to authenticate.
// It implements the authentication repository interface
type JWKS struct {
	set *jwk.Set
}

// New creates a new JWKS struct, fetching the JSON Web Key Set using the URL passed.
func New(url string) (*JWKS, error) {
	set, err := jwk.FetchHTTP(url)
	if err != nil {
		return nil, errors.New("unable to fetch JSON Web Key Set")
	}

	return &JWKS{
		set: set,
	}, nil
}

// GetPublicKey searches inside a JWKS for a key corresponding to the key ID passed
// and, if it's found, it generates the corresponding public RSA key using the key info.
func (jwks *JWKS) GetPublicKey(keyID string) (interface{}, error) {
	if keys := jwks.set.LookupKeyID(keyID); len(keys) == 1 {
		return keys[0].Materialize()
	}

	return nil, errors.New("unable to find key")
}
