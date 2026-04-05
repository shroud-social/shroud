package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type AmbassadorEnvironment struct {
}

type CoreEnvironment struct {
	PostgresUri string `env:"POSTGRES_URI,required"`
}

type DispatcherEnvironment struct {
}

type EchoEnvironment struct {
}

type PersistenceEnvironment struct {
}

type SatelliteEnvironment struct {
	UserAuthSecret   string `env:"USER_AUTH_SECRET,required"`
	UserUploadSecret string `env:"USER_UPLOAD_SECRET,required"`
	RedisUri         string `env:"REDIS_URI,required"`
}

type SeekerEnvironment struct {
}

func Load[T any]() (*T, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	environment := new(T)
	err = env.Parse(environment)
	if err != nil {
		return nil, err
	}

	return environment, nil
}
