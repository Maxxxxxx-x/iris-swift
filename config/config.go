package config

import (
	"errors"
	"os"

	"github.com/Maxxxxxx-x/iris-swift/utils"
	"github.com/goccy/go-yaml"
)

var (
	ErrMissingAppEnv   = errors.New("Missing APP_ENV from Environment")
	ErrMissingAppHost  = errors.New("Missing APP_HOST from Environment")
	ErrMissingAppPort  = errors.New("Missing APP_PORT from Environment")
	ErrMissingSMTPHost = errors.New("Missing SMTP_HOST from Environment")
	ErrMissingSMTPPort = errors.New("Missing SMTP_PORT from Environment")
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

	return config, nil
}
