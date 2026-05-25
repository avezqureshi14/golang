package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "group-1",
	})

	defer reader.Close()

	log.Println("consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("read error:", err)
			continue
		}

		log.Println("received:", string(msg.Value))

		process(msg.Value)
	}
}

func process(data []byte) {
	log.Println("processing:", string(data))
}
