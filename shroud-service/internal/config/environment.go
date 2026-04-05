package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Environment struct {
	UserAuthSecret   string `env:"USER_AUTH_SECRET,required"`
	UserUploadSecret string `env:"USER_UPLOAD_SECRET,required"`
}

func LoadEnvironment() (*Environment, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	environment := &Environment{}
	err = env.Parse(environment)
	if err != nil {
		return nil, err
	}

	return environment, nil
}
