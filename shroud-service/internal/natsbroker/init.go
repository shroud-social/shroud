package natsbroker

import (
	"github.com/nats-io/nats.go"
)

func init() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}

	defer nc.Drain()

}
