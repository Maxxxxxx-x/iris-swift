package config

import (
	"errors"
	"os"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/goccy/go-yaml"
)

var (
	ErrMissingAppEnv                = errors.New("Missing APP_ENV from Environment")
	ErrMissingAppHost               = errors.New("Missing APP_HOST from Environment")
	ErrMissingAppPort               = errors.New("Missing APP_PORT from Environment")
	ErrMissingSMTPHost              = errors.New("Missing SMTP_HOST from Environment")
	ErrMissingSMTPPort              = errors.New("Missing SMTP_PORT from Environment")
	ErrMissingJWTAccessTokenSecret  = errors.New("Missing JWT_ACCESS_TOKEN_SECRET from Environment")
	ErrMissingJWTRefreshTokenSecret = errors.New("Missing JWT_REFRESH_TOKEN_SECRET from Environment")
	ErrMissingJWTAccessTokenTTL     = errors.New("Missing JWT_ACCESS_TOKEN_TTL from Environment")
	ErrMissingJWTRefreshTokenTTL    = errors.New("Missing JWT_REFRESH_TOKEN_TTL from Environment")
)

func getConfigFromEnv() (EnvConfig, error) {
	var config EnvConfig
	var ok bool

	if config.App_Env, ok = os.LookupEnv("APP_ENV"); !ok {
		return config, ErrMissingAppEnv
	}
	if config.App_Host, ok = os.LookupEnv("APP_HOST"); !ok {
		return config, ErrMissingAppHost
	}
	if config.App_Port, ok = os.LookupEnv("APP_PORT"); !ok {
		return config, ErrMissingAppPort
	}
	if config.SMTP_Host, ok = os.LookupEnv("SMTP_HOST"); !ok {
		return config, ErrMissingSMTPHost
	}
	if config.SMTP_Port, ok = os.LookupEnv("SMTP_PORT"); !ok {
		return config, ErrMissingSMTPPort
	}

	return config, nil
}

func readYamlFile(filePath string) ([]byte, error) {
	if err := utils.EnsureFileExists(filePath); err != nil {
		return []byte{}, err
	}
	return os.ReadFile(filePath)
}

func GetAppEnv() string {
	env, ok := os.LookupEnv("APP_ENV")
	if !ok {
		return "dev"
	}
	return env
}

func getJwtConfig() (JWTConfig, error) {
	var config JWTConfig
	var ok bool

	if config.AccessTokenSecret, ok = os.LookupEnv("JWT_ACCESS_TOKEN_SECRET"); !ok {
		return config, ErrMissingJWTAccessTokenSecret
	}

	if config.RefreshTokenSecret, ok = os.LookupEnv("JWT_REFRESH_TOKEN_SECRET"); !ok {
		return config, ErrMissingJWTRefreshTokenSecret
	}

	accessTokenTtlStr, ok := os.LookupEnv("JWT_ACCESS_TOKEN_TTL")
	if !ok {
		return config, ErrMissingJWTAccessTokenTTL
	}
	accessTokenTtl, err := time.ParseDuration(accessTokenTtlStr)
	if err != nil {
		return config, err
	}
	config.AccessTokenTTL = accessTokenTtl

	refreshTokenTtlStr, ok := os.LookupEnv("JWT_REFRESH_TOKEN_TTL")
	if !ok {
		return config, ErrMissingJWTRefreshTokenTTL
	}
	refreshTokenTtl, err := time.ParseDuration(refreshTokenTtlStr)
	if err != nil {
		return config, err
	}
	config.RefreshTokenTTL = refreshTokenTtl

	return config, nil
}

func GetConfig(filePath string) (Config, error) {
	var config Config
	var err error

	if config.Env, err = getConfigFromEnv(); err != nil {
		return config, err
	}

	configBytes, err := readYamlFile(filePath)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		return config, err
	}

	jwtConfig, err := getJwtConfig()
	if err != nil {
		return config, err
	}
	config.JwtConfig = jwtConfig

	return config, nil
}
