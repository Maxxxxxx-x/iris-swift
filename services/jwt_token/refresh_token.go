package token

import (
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/cristalhq/jwt/v5"
)

func GenerateRefreshToken(userId string) (*jwt.Token, error) {
	tokenId, err := utils.GenerateULID()
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to create TokenId while generating refresh token")
		return nil, err
	}

	currentTimestamp := time.Now()

	claims := &RefreshTokenClaims{
		UserId:        userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			Subject:   "refresh-token",
			IssuedAt:  jwt.NewNumericDate(currentTimestamp),
			NotBefore: jwt.NewNumericDate(currentTimestamp),
			ExpiresAt: jwt.NewNumericDate(currentTimestamp.Add(manager.refresh.TTL)),
		},
	}

	token, err := jwt.NewBuilder(manager.refresh.Signer).Build(claims)
	return token, err
}

func VerifyAndParseRefreshToken(rawToken string) (*jwt.Token, error) {
	return manager.ParseAndVerify(rawToken, manager.refresh.Verifier)
}

func VerifyRefreshToken(rawToken string) error {
	return manager.VerifyToken(manager.refresh.Verifier, rawToken)
}


func GetRefreshTokenTTL() time.Duration {
	return manager.refresh.TTL
}
