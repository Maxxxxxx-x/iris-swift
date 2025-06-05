package token

import (
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/cristalhq/jwt/v5"
)

func createSigner(secret string) (jwt.Signer, error) {
	secretBytes := []byte(secret)
	return jwt.NewSignerHS(jwt.HS512, secretBytes)
}

func New(config config.JWTConfig) (ITokenManager, error) {
	accessTokenSigner, err := createSigner(config.AccessTokenSecret)
	if err != nil {
		return nil, err
	}
	refreshTokenSigner, err := createSigner(config.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &TokenManager{
		config:             config,
		accessTokenSigner:  accessTokenSigner,
		refreshTokenSigner: refreshTokenSigner,
	}, nil
}

func (manager *TokenManager) GenerateAccessToken(userId string) (string, string, error) {
}
