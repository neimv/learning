package main

import (
	"fmt"
	"log"
	"net/http"
	"platzi-cqrs/events"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf("%v", err)
	}

	hub := NewHub()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))

	if err != nil {
		log.Fatal(err)
	}

	err = n.OnCreateFeed(func(cfm events.CreatedFeedMessage) {
		hub.Broadcast(newCreatedFeedMessage(cfm.ID, cfm.Title, cfm.Description, cfm.CreatedAt), nil)
	})

	if err != nil {
		log.Fatal(err)
	}

	events.SetEventStore(n)

	defer events.Close()

	go hub.Run()

	http.HandleFunc("/ws", hub.HandleWebSocket)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
