package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/goccy/go-yaml"
)

var (
	ErrMissingEnvVar = func(name string) error {
		return errors.New(fmt.Sprintf("Missing %s from Environment", name))
	}
)

func getFromEnv(name string) (string, error) {
	config, ok := os.LookupEnv(name)
	if !ok {
		return "", ErrMissingEnvVar(name)
	}
	return config, nil
}

func getAppConfig() (AppConfig, error) {
	host, err := getFromEnv("APP_HOST")
	if err != nil {
		return AppConfig{}, err
	}

	port, err := getFromEnv("APP_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Host: host,
		Port: port,
	}, nil
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

func getTokenConfig(secretEnv, ttlEnv string) (JWTTokenConfig, error) {
	var config JWTTokenConfig
	secret, ok := os.LookupEnv(secretEnv)
	if !ok {
		return config, ErrMissingEnvVar(secretEnv)
	}

	rawTTL, ok := os.LookupEnv(ttlEnv)
	if !ok {
		return config, ErrMissingEnvVar(ttlEnv)
	}

	ttl, err := time.ParseDuration(rawTTL)
	if err != nil {
		return config, err
	}

	config.Secret = secret
	config.TTL = ttl
	return config, nil
}

func getJwtConfig() (JWTConfig, error) {

	accessToken, err := getTokenConfig("JWT_ACCESS_TOKEN_SECRET", "JWT_ACCESS_TOKEN_TTL")
	if err != nil {
		return JWTConfig{}, err
	}

	refreshToken, err := getTokenConfig("JWT_REFRESH_TOKEN_SECRET", "JWT_REFRESH_TOKEN_TTL")
	if err != nil {
		return JWTConfig{}, err
	}

	verifyEmailToken, err := getTokenConfig("JWT_VERIFY_EMAIL_TOKEN_SECRET", "JWT_VERIFY_EMAIL_TOKEN_TTL")
	if err != nil {
		return JWTConfig{}, err
	}

	resetPasswordToken, err := getTokenConfig("JWT_RESET_PASSWORD_TOKEN_SECRET", "JWT_RESET_PASSWORD_TOKEN_TTL")
	if err != nil {
		return JWTConfig{}, err
	}

	return JWTConfig{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		VerifyEmailToken:   verifyEmailToken,
		ResetPasswordToken: resetPasswordToken,
	}, nil
}

func GetConfig(filePath string) (Config, error) {
	var config Config
	var err error

	appConfig, err := getAppConfig()
	if err != nil {
		return config, err
	}
	config.App = appConfig

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
