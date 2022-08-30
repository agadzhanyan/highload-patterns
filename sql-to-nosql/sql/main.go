package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	log.SetOutput(os.Stdout)

	connStr := "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		uid := uuid.New().String()
		blob, _ := json.Marshal(randSeq(100))

		err := db.QueryRowContext(
			request.Context(),
			`INSERT INTO activities(id, user_id, timestamp, data) VALUES($1, $2, $3, $4)`,
			uid,
			uuid.New().String(),
			time.Now(),
			string(blob),
		).Err()
		if err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.Write([]byte("failed!" + err.Error()))
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte("saved! " + uid))
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
