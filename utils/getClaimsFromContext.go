package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrNoClaimsFound   = errors.New("No claims found in context")
	ErrWrongClaimsType = errors.New("Claims in context are of the wrong type")
)

func GetClaimsFromContext[T any](ctx echo.Context, key string) (*T, error) {
	raw := ctx.Get(key)
	if raw == nil {
		return nil, ErrNoClaimsFound
	}
	claims, ok := raw.(*T)
	if !ok {
		return nil, ErrWrongClaimsType
	}

	return claims, nil
}
