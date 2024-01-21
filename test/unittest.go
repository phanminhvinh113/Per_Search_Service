package test

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ConcurrentHandler handles concurrent requests using Goroutines
func ConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	// Use a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Number of concurrent Goroutines to spawn
	numRequests := 10

	// Increment the WaitGroup counter
	wg.Add(numRequests)

	// Handle each request concurrently
	for i := 0; i < numRequests; i++ {
		go func(index int) {
			// Decrement the WaitGroup counter when the Goroutine completes
			defer wg.Done()

			// Simulate some processing time
			time.Sleep(time.Second)

			// Respond to the client
			response := fmt.Sprintf("Response from Goroutine %d", index)
			fmt.Fprintln(w, response)
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
}
