package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Movie struct {
	Title string `json:"Title"`
}

type CachedPopularItems struct {
	lock   sync.RWMutex
	Movies []Movie
}

func main() {
	ctx := context.Background()

	// initializing cache and fill
	cache := CachedPopularItems{}
	cache.Movies = getPopularMoviesFromDB()
	go func() {
		timer := time.NewTicker(1 * time.Second)
		defer timer.Stop()

		// initializing background job
		for {
			select {
			// refreshing cache
			case <-timer.C:
				movies := getPopularMoviesFromDB()

				// updating cache struct
				cache.lock.Lock()
				cache.Movies = movies
				cache.lock.Unlock()

			// app is terminating
			case <-ctx.Done():
				break
			}
		}
	}()

	http.HandleFunc("/getPopularMovies", func(writer http.ResponseWriter, request *http.Request) {
		cache.lock.RLock()
		movies := cache.Movies
		cache.lock.RUnlock()

		bytes, _ := json.Marshal(movies)

		writer.Header().Add("Content-Type", "application/json")
		writer.Write(bytes)
	})

	_ = http.ListenAndServe(":8890", nil)
}

// Getting from DB
func getPopularMoviesFromDB() []Movie {
	// simulation request to database with latency
	time.Sleep(5 * time.Second)

	return []Movie{{Title: "Avatar"}, {Title: "I Am Legend"}, {Title: "The Wolf of Wall Street"}}
}
