package token

import (
	"encoding/json"

	"github.com/cristalhq/jwt/v5"
)

func (manager *TokenManager) ParseNoVerify(rawToken string) (*jwt.Token, error) {
	return jwt.ParseNoVerify([]byte(rawToken))
}

func (manager *TokenManager) ParseAndVerify(rawToken string, verifier jwt.Verifier) (*jwt.Token, error) {
	return jwt.Parse([]byte(rawToken), verifier)
}

func (manager *TokenManager) VerifyToken(verifier jwt.Verifier, rawToken string) error {
	token, err := manager.ParseNoVerify(rawToken)
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to parse token")
		return err
	}

	return verifier.Verify(token)
}

func GetRegisteredClaims(rawToken string) (jwt.RegisteredClaims, error) {
	token, err := manager.ParseNoVerify(rawToken)
	if err != nil {
		manager.logger.Error().Err(err).Msg("Failed to parse token")
		return jwt.RegisteredClaims{}, err
	}
	var claims jwt.RegisteredClaims
	if err := json.Unmarshal(token.Claims(), &claims); err != nil {
		manager.logger.Error().Err(err).Msg("Failed to get registered claims from token")
		return jwt.RegisteredClaims{}, err
	}
	return claims, nil
}

func VerifyAndGetClaims(rawToken string, verifier jwt.Verifier, claims *any) error {
	return jwt.ParseClaims([]byte(rawToken), verifier, claims)
}
