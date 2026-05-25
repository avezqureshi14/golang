package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	defer writer.Close()

	for i := 0; i < 5; i++ {
		msg := kafka.Message{
			Key:   []byte("key"),
			Value: []byte("task number " + string(rune(i+'0'))),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatal("write error:", err)
		}

		log.Println("sent:", string(msg.Value))
		time.Sleep(2 * time.Second)
	}
}
