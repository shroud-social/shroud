package main

import (
	"services/internal/api"
	"services/internal/config"
)

func main() {
	environment, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	api.LoadAuthSecret(environment.UserAuthSecret)
	api.LoadUploadSecret(environment.UserUploadSecret)

	api.StartRouter()
}
