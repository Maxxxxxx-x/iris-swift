package token

import (
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/cristalhq/jwt/v5"
)

func GenerateVerifyEmailToken(email string) (*jwt.Token, error) {
	tokenId, err := utils.GenerateULID()
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to create TokenId while generating verify email token")
		return nil, err
	}

	currentTimestamp := time.Now()

	claims := &VerifyEmailTokenClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			Subject:   "verify-email;",
			IssuedAt:  jwt.NewNumericDate(currentTimestamp),
			NotBefore: jwt.NewNumericDate(currentTimestamp),
			ExpiresAt: jwt.NewNumericDate(currentTimestamp.Add(manager.verifyEmail.TTL)),
		},
	}

	token, err := jwt.NewBuilder(manager.verifyEmail.Signer).Build(claims)
	return token, err
}

func VerifyAndParseVerifyEmailToken(rawToken string) (*jwt.Token, error) {
	return manager.ParseAndVerify(rawToken, manager.verifyEmail.Verifier)
}

func VerifyVerifyEmailToken(rawToken string) error {
	return manager.VerifyToken(manager.verifyEmail.Verifier, rawToken)
}

func GetVerifyEmailTokenTTL() time.Duration {
	return manager.verifyEmail.TTL
}
