package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
)

// docker run --network=one-instance_backend --rm skandyla/wrk -t1 -c1 -d5s http://app:8890/handle
// docker run --network=one-instance_backend --rm skandyla/wrk -t3 -c3 -d5s http://app:8890/handle

// docker run --network=async_backend --rm skandyla/wrk -t3 -c3 -d5s http://app:8890/handle
// docker run --network=async_backend --rm skandyla/wrk -t1 -c1 -d5s http://app:8890/handle
// docker-compose up --scale app=5
// docker-compose up -d --build --force-recreate
// docker-compose up --build --force-recreate
// wrk -t3 -c3 -d5s http://localhost:8891/handle
func main() {
	log.SetOutput(os.Stdout)

	// kafka topic
	topic := "my-topic"
	partition := 0

	// create kafka connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", "broker:29092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	http.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		// cpu intensive work
		for i := 0; i <= 1000; i++ {
			_, _ = json.Marshal(randSeq(10))
		}

		// produce event to kafka
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(randSeq(100))},
		)
		if err != nil {
			log.Println(err)
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte("done!"))
	})

	_ = http.ListenAndServe(":8890", nil)
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
