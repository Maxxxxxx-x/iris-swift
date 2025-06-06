package authHandler

import (
	"net/http"
	"time"

	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/cristalhq/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	ACCESS_TOKEN_NAME = "access-token"
	REFRESH_TOKEN_NAME = "refresh-token"
)

func (handler *AuthHandler) setAccessTokenToCookie(ctx echo.Context, accessToken *jwt.Token) {
	currentTimestamp := time.Now()
	ctx.SetCookie(&http.Cookie{
		Name:     ACCESS_TOKEN_NAME,
		Value:    accessToken.String(),
		Path:     "/",
		Expires:  currentTimestamp.Add(token.GetAccessTokenTTL()),
		MaxAge:   int(currentTimestamp.Add(token.GetAccessTokenTTL()).Unix()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func (handler *AuthHandler) removeAccessTokenFromCookie(ctx echo.Context) {
	exist := ctx.Get(ACCESS_TOKEN_NAME)
	if exist == nil {
		return
	}
	ctx.SetCookie(&http.Cookie{
		Name: ACCESS_TOKEN_NAME,
		Value: "",
		Path: "/",
		Expires: time.Unix(0, 0),
		MaxAge: 0,
		HttpOnly: true,
	})
}

func (handler *AuthHandler) setRefreshTokenToCookie(ctx echo.Context, refreshToken *jwt.Token) {
	currentTimestamp := time.Now()

	ctx.SetCookie(&http.Cookie{
		Name:     REFRESH_TOKEN_NAME,
		Value:    refreshToken.String(),
		Path:     "/api/v1/auth/refresh",
		Expires:  currentTimestamp.Add(token.GetRefreshTokenTTL()),
		MaxAge:   int(currentTimestamp.Add(token.GetRefreshTokenTTL()).Unix()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func (handler *AuthHandler) removeRefreshTokenFromCookie(ctx echo.Context) {
	exist := ctx.Get(REFRESH_TOKEN_NAME)
	if exist == nil {
		return
	}
	ctx.SetCookie(&http.Cookie{
		Name: REFRESH_TOKEN_NAME,
		Value: "",
		Path: "/",
		Expires: time.Unix(0, 0),
		MaxAge: 0,
		HttpOnly: true,
	})
}
