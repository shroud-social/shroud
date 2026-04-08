package pubsub

import (
	"github.com/nats-io/nats.go"
)

var Connection *nats.Conn

func Connect(uri string) error {
	Conn, err := nats.Connect(uri)
	Connection = Conn
	return err
}
