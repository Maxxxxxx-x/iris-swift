package token

import (
	"errors"
	"fmt"

	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/Maxxxxxx-x/iris-swift/logger"
	"github.com/cristalhq/jwt/v5"
)

var (
	ErrInitializationFailed = func(name string, err error) error {
		return errors.New(fmt.Sprint("Failed to initialize signer verifier pair for %s: %v\n", name, err.Error()))
	}
)

const (
	TOKEN_ISSUER = "iris-auth"
)

var manager *TokenManager

func newSignerVerifier(config config.JWTTokenConfig) (SignerVerifier, error) {
	secretBytes := []byte(config.Secret)
	signer, err := jwt.NewSignerHS(jwt.HS512, secretBytes)
	if err != nil {
		return SignerVerifier{}, err
	}

	verifier, err := jwt.NewVerifierHS(jwt.HS512, secretBytes)
	if err != nil {
		return SignerVerifier{}, err
	}

	return SignerVerifier{
		Signer:   signer,
		Verifier: verifier,
		TTL:      config.TTL,
	}, nil
}

func Init(cfg config.JWTConfig) error {
	logger := logger.NewLogger("jwt-manager")
	access, err := newSignerVerifier(cfg.AccessToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create signer-verifier pair for AccessToken")
		return ErrInitializationFailed("access", err)
	}

	refresh, err := newSignerVerifier(cfg.RefreshToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create signer-verifier pair fori RefreshToken")
		return ErrInitializationFailed("refresh", err)
	}

	verifyEmail, err := newSignerVerifier(cfg.VerifyEmailToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create signer-verifier pair for VerifyEmailToken")
		return ErrInitializationFailed("verify email", err)
	}

	resetPassword, err := newSignerVerifier(cfg.ResetPasswordToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create signer-verifier pair for ResetPasswordToken")
		return ErrInitializationFailed("reset password", err)
	}

	manager = &TokenManager{
		access:        access,
		refresh:       refresh,
		verifyEmail:   verifyEmail,
		resetPassword: resetPassword,
		logger:        logger,
	}

	return nil
}
