package token

import (
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/rs/zerolog"
)

type ITokenManager interface {
	GenerateAccessToken(userId, role string) (*jwt.Token, string, error)
	GenerateRefreshToken(userId, accessTokenId string) (*jwt.Token, error)
	GenerateVerifyEmailToken(email string) (*jwt.Token, error)
	GenerateResetPasswordToken(email string) (*jwt.Token, error)
	VerifyAccessToken(rawToken string) (*jwt.Token, error)
	VerifyRefreshToken(rawToken string) (*jwt.Token, error)
	VerifyVerifyEmailToken(rawToken string) (*jwt.Token, error)
	VerifyResetPasswordToken(rawToken string) (*jwt.Token, error)
}

type TokenManager struct {
	access        SignerVerifier
	refresh       SignerVerifier
	verifyEmail   SignerVerifier
	resetPassword SignerVerifier
	logger        zerolog.Logger
}

type SignerVerifier struct {
	Signer   jwt.Signer
	Verifier jwt.Verifier
	TTL      time.Duration
}

type AccessTokenClaims struct {
	UserId      string `json:"user_id"`
	AccountType string `json:"account_type"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserId        string `json:"user_id"`
	jwt.RegisteredClaims
}

type VerifyEmailTokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type ResetPasswordTokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
