package config

import (
	"github.com/caarlos0/env"
)

type AmbassadorEnvironment struct {
	NatsUri string `env:"NATS_URI,required"`
}

type CoreEnvironment struct {
	NatsUri     string `env:"NATS_URI,required"`
	PostgresUri string `env:"POSTGRES_URI,required"`
}

type DispatcherEnvironment struct {
	NatsUri string `env:"NATS_URI,required"`
}

type EchoEnvironment struct {
	NatsUri string `env:"NATS_URI,required"`
}

type ScribeEnvironment struct {
	NatsUri   string `env:"NATS_URI,required"`
	ScyllaUri string `env:"SCYLLA_URI,required"`
}

type SatelliteEnvironment struct {
	NatsUri          string `env:"NATS_URI,required"`
	UserAuthSecret   string `env:"USER_AUTH_SECRET,required"`
	UserUploadSecret string `env:"USER_UPLOAD_SECRET,required"`
	UserUploadUri    string `env:"USER_UPLOAD_URI,required"`
	RedisUri         string `env:"REDIS_URI,required"`
	RedisPassword    string `env:"REDIS_PASSWORD"`
}

type SeekerEnvironment struct {
	NatsUri string `env:"NATS_URI,required"`
}

func Load[T any]() (*T, error) {
	environment := new(T)
	err := env.Parse(environment)
	if err != nil {
		return nil, err
	}
	return environment, nil
}
