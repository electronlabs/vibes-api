package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/electronlabs/vibes-api/domain/auth/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthService(t *testing.T) {
	Convey("Valid Token", t, func() {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatal("Could not generate RSA key pair")
		}

		audience := "test-audience"
		issuer := "test-issuer"
		keyID := "test-key-id"
		repository := &mocks.AuthRepository{}
		repository.On("GetPublicKey", mock.Anything).Return(&privateKey.PublicKey, nil)

		claims := &jwt.StandardClaims{
			Audience:  audience,
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
		token.Header["kid"] = keyID

		signedToken, err := token.SignedString(privateKey)
		if err != nil {
			t.Fatal(err)
		}

		service := NewService(repository, &Config{Audience: audience, Issuer: issuer})

		result, err := service.CheckJWT(signedToken)

		So(err, ShouldBeNil)
		So(result.Raw, ShouldEqual, signedToken)
		So(result.Valid, ShouldBeTrue)
	})
}
