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
	Convey("Token Verification", t, func() {
		Convey("Successful Authentication", func() {
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

		Convey("Fails if any of the claims is not correct", func() {
			Convey("Invalid token audience", func() {
				privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
				if err != nil {
					t.Fatal("Could not generate RSA key pair")
				}

				svcAudience := "test-audience"
				tokenAudience := "some-other-audience"
				issuer := "test-issuer"
				keyID := "test-key-id"

				repository := &mocks.AuthRepository{}
				repository.On("GetPublicKey", mock.Anything).Return(&privateKey.PublicKey, nil)
				service := NewService(repository, &Config{Audience: svcAudience, Issuer: issuer})

				claims := &jwt.StandardClaims{
					Audience:  tokenAudience,
					Issuer:    issuer,
					ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
				token.Header["kid"] = keyID

				signedToken, err := token.SignedString(privateKey)
				if err != nil {
					t.Fatal(err)
				}

				result, err := service.CheckJWT(signedToken)
				So(err, ShouldBeError)
				So(err.Error(), ShouldEqual, invalidAudience)
				So(result, ShouldBeNil)
			})

			Convey("Invalid token issuer", func() {
				privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
				if err != nil {
					t.Fatal("Could not generate RSA key pair")
				}

				audience := "test-audience"
				svcIssuer := "test-issuer"
				tokenIssuer := "some-other-issuer"
				keyID := "test-key-id"

				repository := &mocks.AuthRepository{}
				repository.On("GetPublicKey", mock.Anything).Return(&privateKey.PublicKey, nil)
				service := NewService(repository, &Config{Audience: audience, Issuer: svcIssuer})

				claims := &jwt.StandardClaims{
					Audience:  audience,
					Issuer:    tokenIssuer,
					ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
				token.Header["kid"] = keyID

				signedToken, err := token.SignedString(privateKey)
				if err != nil {
					t.Fatal(err)
				}

				result, err := service.CheckJWT(signedToken)
				So(err, ShouldBeError)
				So(err.Error(), ShouldEqual, invalidIssuer)
				So(result, ShouldBeNil)
			})

			Convey("Expiration date has passed", func() {
				privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
				if err != nil {
					t.Fatal("Could not generate RSA key pair")
				}

				audience := "test-audience"
				issuer := "test-issuer"
				keyID := "test-key-id"

				repository := &mocks.AuthRepository{}
				repository.On("GetPublicKey", mock.Anything).Return(&privateKey.PublicKey, nil)
				service := NewService(repository, &Config{Audience: audience, Issuer: issuer})

				claims := &jwt.StandardClaims{
					Audience:  audience,
					Issuer:    issuer,
					ExpiresAt: time.Now().Add(time.Minute * (-1)).Unix(),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
				token.Header["kid"] = keyID

				signedToken, err := token.SignedString(privateKey)
				if err != nil {
					t.Fatal(err)
				}

				result, err := service.CheckJWT(signedToken)

				// Error message returned from go-jwt lib, just check if it's not nil
				So(err, ShouldBeError)
				So(result, ShouldBeNil)
			})
		})

		Convey("Fails if signature verification returns an error", func() {
			privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				t.Fatal("Could not generate RSA key pair")
			}

			anotherPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				t.Fatal("Could not generate RSA key pair")
			}

			audience := "test-audience"
			issuer := "test-issuer"
			keyID := "test-key-id"

			repository := &mocks.AuthRepository{}
			repository.On("GetPublicKey", mock.Anything).Return(&anotherPrivateKey.PublicKey, nil)
			service := NewService(repository, &Config{Audience: audience, Issuer: issuer})

			claims := &jwt.StandardClaims{
				Audience:  audience,
				Issuer:    issuer,
				ExpiresAt: time.Now().Add(time.Minute * (-1)).Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
			token.Header["kid"] = keyID

			signedToken, err := token.SignedString(privateKey)
			if err != nil {
				t.Fatal(err)
			}

			result, err := service.CheckJWT(signedToken)

			// Error message returned from go-jwt lib, just check if it's not nil
			So(err, ShouldBeError)
			So(result, ShouldBeNil)
		})
	})
}
