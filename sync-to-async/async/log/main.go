package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	log.SetOutput(os.Stdout)

	// kafka topic
	topic := "my-topic"
	partition := 0

	// create kafka connection
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"broker:29092"},
		Topic:     topic,
		Partition: partition,
	})

	for {
		// read message from kafka
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			break
		}

		// cpu intensive work
		for i := 0; i <= 1000; i++ {
			_, _ = json.Marshal(randSeq(50))
		}

		log.Print(m.Offset, string(m.Key), string(m.Value))
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Random string generator
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
