package token

import (
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/cristalhq/jwt/v5"
)

func GenerateResetPasswordToken(email string) (*jwt.Token, error) {
	tokenId, err := utils.GenerateULID()
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to create TokenId while generating reset password token")
		return nil, err
	}

	currentTimestamp := time.Now()

	claims := &ResetPasswordTokenClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			Subject:   "reset-password",
			IssuedAt:  jwt.NewNumericDate(currentTimestamp),
			NotBefore: jwt.NewNumericDate(currentTimestamp),
			ExpiresAt: jwt.NewNumericDate(currentTimestamp.Add(manager.resetPassword.TTL)),
		},
	}

	token, err := jwt.NewBuilder(manager.resetPassword.Signer).Build(claims)
	return token, err
}


func VerifyAndParseResetPasswordToken(rawToken string) (*jwt.Token, error) {
	return manager.ParseAndVerify(rawToken, manager.resetPassword.Verifier)
}

func VerifyResetPasswordToken(rawToken string) error {
	return manager.VerifyToken(manager.resetPassword.Verifier, rawToken)
}

func GetResetPasswordTokenTTL() time.Duration {
	return manager.resetPassword.TTL
}
