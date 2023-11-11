package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// This script is to test functionality of the API of GetUserRecord v1 when applying HTTP caching and no caching
// This will test in 3 cases
// 1. GET 1 record for multiple times concurrently
// 2. GET multiple records for multiple times concurrently (small size)
// 3. GET multiple records for multiple times concurrently (large size)

const (
	httpRequestTest = "http://localhost:9000/api/v1/GetUserRecord/1"
)

func main() {
	// Start measuring time
	startTime := time.Now()

	response, err := http.Get(httpRequestTest)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	elapsedTime := time.Since(startTime)

	log.Println(response)

	log.Println("Time for request: " + GetTimeForRequest(elapsedTime))

}

// Calculate time from sending HTTP request to receiving response
func GetTimeForRequest(elapsedTime time.Duration) string {
	return elapsedTime.String()
}
