package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	log.SetOutput(os.Stdout)

	cluster := gocql.NewCluster("cassandra", "cassandra2", "cassandra3")
	cluster.Keyspace = "app"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	http.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		uid := gocql.UUIDFromTime(time.Now())
		blob, _ := json.Marshal(randSeq(100))
		if err := session.Query(
			`INSERT INTO activities (id, user_id, timestamp, data) VALUES (?, ?, ?, ?)`,
			uid,
			gocql.UUIDFromTime(time.Now()),
			time.Now().Unix(),
			blob,
		).Exec(); err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.Write([]byte("failed!" + err.Error()))
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte("saved! " + uid.String()))
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
