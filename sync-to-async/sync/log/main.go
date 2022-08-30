package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/sendEvent", func(writer http.ResponseWriter, request *http.Request) {
		// cpu intensive work
		for i := 0; i <= 1000; i++ {
			_, _ = json.Marshal(randSeq(50))
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte("event is saved!"))
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
