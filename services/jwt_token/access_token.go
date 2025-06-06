package token

import (
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/cristalhq/jwt/v5"
)

func GenerateAccessToken(userId, role string) (*jwt.Token, error) {
	tokenId, err := utils.GenerateULID()
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to create TokenId while generating access token")
		return nil, err
	}

	currentTimestamp := time.Now()

	claims := &AccessTokenClaims{
		UserId:      userId,
		AccountType: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			Subject:   "access-token",
			IssuedAt:  jwt.NewNumericDate(currentTimestamp),
			NotBefore: jwt.NewNumericDate(currentTimestamp),
			ExpiresAt: jwt.NewNumericDate(currentTimestamp.Add(manager.access.TTL)),
		},
	}

	token, err := jwt.NewBuilder(manager.access.Signer).Build(claims)
	return token, err
}

func VerifyAndParseAccessToken(rawToken string) (*jwt.Token, error) {
	return manager.ParseAndVerify(rawToken, manager.access.Verifier)
}

func VerifyAccessToken(rawToken string) error {
	return manager.VerifyToken(manager.access.Verifier, rawToken)
}

func GetAccessTokenTTL() time.Duration {
	return manager.access.TTL
}
