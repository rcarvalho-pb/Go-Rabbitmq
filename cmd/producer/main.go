package main

import (
	"github.com/rcarvalho-pb/goutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello Wordl!", "amq.direct")
}