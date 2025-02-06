package config

import (
	"fmt"
	"os"
)

const (
	_MYSQL_ROOT_PASSWORD = "MYSQL_ROOT_PASSWORD"
	_MYSQL_DATABASE      = "MYSQL_DATABASE"
	_MYSQL_USER          = "MYSQL_USER"
	_MYSQL_PASSWORD      = "MYSQL_PASSWORD"
	_MYSQL_HOST          = "MYSQL_HOST"
	_MYSQL_PORT          = "MYSQL_PORT"
)

var (
	requiredDBEnv = []string{
		_MYSQL_ROOT_PASSWORD,
		_MYSQL_DATABASE,
		_MYSQL_USER,
		_MYSQL_PASSWORD,
	}
	requiredAppEnv = []string{
		_MYSQL_ROOT_PASSWORD,
		_MYSQL_DATABASE,
		_MYSQL_USER,
		_MYSQL_PASSWORD,
	}
)

type Config struct {
	MYSQL_ROOT_PASSWORD string
	MYSQL_DATABASE      string
	MYSQL_USER          string
	MYSQL_PASSWORD      string
	MYSQL_HOST          string
	MYSQL_PORT          string
}

func NewConfig() (*Config, error) {
	if err := ValidateDBEnv(); err != nil {
		return nil, err
	}
	if err := ValidateAppEnv(); err != nil {
		return nil, err
	}

	return &Config{
		MYSQL_ROOT_PASSWORD: os.Getenv(_MYSQL_ROOT_PASSWORD),
		MYSQL_DATABASE:      os.Getenv(_MYSQL_DATABASE),
		MYSQL_USER:          os.Getenv(_MYSQL_USER),
		MYSQL_PASSWORD:      os.Getenv(_MYSQL_PASSWORD),
		MYSQL_HOST:          os.Getenv(_MYSQL_HOST),
		MYSQL_PORT:          os.Getenv(_MYSQL_PORT),
	}, nil
}

func ValidateDBEnv() error {
	for _, env := range requiredDBEnv {
		if os.Getenv(env) == "" {
			return fmt.Errorf("missing required env %s", env)
		}
	}
	return nil
}

func ValidateAppEnv() error {
	for _, env := range requiredAppEnv {
		if os.Getenv(env) == "" {
			return fmt.Errorf("missing required env %s", env)
		}
	}
	return nil
}
