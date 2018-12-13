package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/electronlabs/vibes-api/domain/actions"
	actionsRoutes "github.com/electronlabs/vibes-api/router/actions"
	healthRoutes "github.com/electronlabs/vibes-api/router/actions"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/electronlabs/vibes-api/config"
	"github.com/gin-gonic/gin"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

var configuration = config.NewConfig()

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(configuration.JWKSURL)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := configuration.Auth0APIAudience
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("Invalid audience.")
		}
		// Verify 'iss' claim
		iss := configuration.Auth0NativeIssuer
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
	}
}

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(actionsSvc *actions.Service) http.Handler {
	router := gin.Default()

	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory(actionsSvc)(healthGroup)

	api := router.Group("/api")
	actionsGroup := api.Group("/actions")
	actionsGroup.Use(checkJWT())
	actionsRoutes.NewRoutesFactory(actionsSvc)(actionsGroup)
	return router
}
