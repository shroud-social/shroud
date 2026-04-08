package main

import (
	"log"
	"os"
	"os/signal"
	"services/internal/api/service"
	"services/internal/comm/pubsub"
	"services/internal/config"
	"syscall"
)

func main() {
	environment, err := config.Load[config.SeekerEnvironment]()
	if err != nil {
		panic(err)
	}

	err = pubsub.Connect(environment.NatsUri)
	if err != nil {
		panic(err)
	}
	err = service.Subscribe(service.SubscriptionsCore)
	if err != nil {
		panic(err)
	}
	log.Println("Shroud Core service started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	log.Println("Shroud Core service shutting down...")
	err = pubsub.Connection.Drain()
	if err != nil {
		panic(err)
	}
	pubsub.Connection.Close()
	log.Println("Shroud Core service shut down")
}
