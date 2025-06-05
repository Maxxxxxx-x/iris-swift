package token

import (
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/cristalhq/jwt/v5"
)

type ITokenManager interface {

}

type TokenManager struct {
	config config.JWTConfig
	accessTokenSigner jwt.Signer
	refreshTokenSigner jwt.Signer
}
