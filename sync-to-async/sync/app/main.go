package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		// cpu intensive work
		for i := 0; i <= 1000; i++ {
			_, _ = json.Marshal(randSeq(10))
		}

		// send log
		resp, err := http.DefaultClient.Post(
			"http://logs:8890/sendEvent",
			"application/json",
			bytes.NewReader([]byte(randSeq(100))),
		)
		if err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.Write([]byte("failed!" + err.Error()))
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			writer.Header().Add("Content-Type", "application/json")
			writer.Write([]byte("failed!" + err.Error()))
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte("done!" + string(body)))
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
