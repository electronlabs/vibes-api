package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"

	authServiceMocks "github.com/electronlabs/vibes-api/domain/auth/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Convey("Auth Middleware", t, func() {
		Convey("Authorization Header", func() {
			Convey("Fails with 400 if authorization header format is not Bearer {token}", func() {
				Convey("No token prefix", func() {
					authSvc := &authServiceMocks.AuthService{}
					authSvc.On("CheckJWT", mock.Anything).Return(&jwt.Token{}, nil)

					rec := httptest.NewRecorder()
					router := gin.Default()
					router.Use(CheckJWT(authSvc))
					router.GET("/", func(c *gin.Context) {
						c.String(http.StatusOK, "ok")
					})

					req := &http.Request{
						Header: http.Header{
							"Authorization": {"xxx.yyy.zzz"},
						},
						URL: &url.URL{},
					}

					router.ServeHTTP(rec, req)

					So(rec.Code, ShouldEqual, 400)
				})

				Convey("Prefix is not Bearer", func() {
					authSvc := &authServiceMocks.AuthService{}
					authSvc.On("CheckJWT", mock.Anything).Return(&jwt.Token{}, nil)

					rec := httptest.NewRecorder()
					router := gin.Default()
					router.Use(CheckJWT(authSvc))
					router.GET("/", func(c *gin.Context) {
						c.String(http.StatusOK, "ok")
					})

					req := &http.Request{
						Header: http.Header{
							"Authorization": {"Foo xxx.yyy.zzz"},
						},
						URL: &url.URL{},
					}

					router.ServeHTTP(rec, req)

					So(rec.Code, ShouldEqual, 400)
				})
			})
		})

		Convey("Token validation", func() {
			Convey("Fails with 401 when authorization returns an error in token validation", func() {
				authSvc := &authServiceMocks.AuthService{}
				authSvc.On("CheckJWT", mock.Anything).Return(nil, errors.New("token validation error"))

				rec := httptest.NewRecorder()
				router := gin.Default()
				router.Use(CheckJWT(authSvc))
				router.GET("/", func(c *gin.Context) {
					c.String(http.StatusOK, "ok")
				})

				req := &http.Request{
					Header: http.Header{
						"Authorization": {"Bearer xxx.yyy.zzz"},
					},
					URL: &url.URL{},
				}

				router.ServeHTTP(rec, req)

				So(rec.Code, ShouldEqual, 401)
			})

			Convey("Writes logged user in context and proceeds to next request handler", func() {
				authSvc := &authServiceMocks.AuthService{}
				token := &jwt.Token{}
				authSvc.On("CheckJWT", mock.Anything).Return(token, nil)

				rec := httptest.NewRecorder()
				ctx, router := gin.CreateTestContext(rec)
				router.Use(CheckJWT(authSvc))
				router.GET("/", func(c *gin.Context) {
					c.String(http.StatusOK, "ok")
				})

				ctx.Request = &http.Request{
					Header: http.Header{
						"Authorization": {"Bearer xxx.yyy.zzz"},
					},
					URL: &url.URL{},
				}

				user, ok := ctx.Get("user")

				So(rec.Code, ShouldEqual, 200)
				So(ok, ShouldEqual, true)
				So(user, ShouldPointTo, token)
			})
		})
	})
}
