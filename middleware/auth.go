package middleware

import (
	"errors"
	"net/http"

	token "github.com/Maxxxxxx-x/iris-swift/services/jwt_token"
	"github.com/labstack/echo/v4"
)

func AccessTokenRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("access-token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing access token")
			}
			return echo.NewHTTPError(http.StatusBadRequest, "missing data in request")
		}

		claims, err := token.VerifyAndParseAccessToken(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired access token")
		}

		ctx.Set("access-token-claims", claims)

		return next(ctx)
	}
}

func RefreshTokenRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("refresh-token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
			}
			return echo.NewHTTPError(http.StatusBadRequest, "missing data in request")
		}

		claims, err := token.VerifyAndParseRefreshToken(cookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired access token")
		}

		ctx.Set("refresh-token-claims", claims)

		return next(ctx)
	}
}
