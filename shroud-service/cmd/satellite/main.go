package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	userapi "services/internal/api/user"
	"services/internal/api/user/v1"
	"services/internal/comm/pubsub"
	"services/internal/config"
	"services/internal/database/redisdb"
	"syscall"
	"time"
)

func main() {
	environment, err := config.Load[config.SatelliteEnvironment]()
	if err != nil {
		panic(err)
	}

	redisdb.LoadConfig(environment.RedisUri, environment.RedisPassword)

	err = pubsub.Connect(environment.NatsUri)
	if err != nil {
		panic(err)
	}

	v1.LoadAuthSecret(environment.UserAuthSecret)
	v1.LoadUploadConf(environment.UserUploadSecret, environment.UserUploadUri)
	srv := userapi.SetupRouters()
	log.Println("Shroud Core service started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	log.Println("Shroud Satellite service shutting down...")

	log.Println("Stopping HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP server stopped")

	log.Println("Draining NATS connection...")
	err = pubsub.Connection.Drain()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("NATS connection stopped")

	log.Println("Closing NATS connection...")
	pubsub.Connection.Close()
	log.Println("NATS connection closed")

	log.Println("Shroud Satellite service shut down")
}
