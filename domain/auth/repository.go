package auth

// AuthRepository interface defines the authentication data source interface.
type AuthRepository interface {
	GetPublicKey(keyID string) (interface{}, error)
}
