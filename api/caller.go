package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	maxRetries = 3
)

var (
	totalLatency              = 0
	numberOfSuccessfulRequest = 0
)

func MakeCallToApi(endpoint string, cacheEnable bool, onlyOneRequest bool, numberOfRequest int) error {
	// Set transport
	transport := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 10000,
	}

	// Set client
	client := &http.Client{
		Transport: transport,
	}

	startTime := time.Now()

	for i := 0; i < maxRetries; i++ {
		resp, err := client.Get(endpoint)
		if err == nil {
			// Process the response.
			numberOfSuccessfulRequest++
			defer resp.Body.Close()
			break
		}
		log.Printf("Error: %v. Retrying...", err)
	}

	latency := time.Since(startTime).Microseconds()

	totalLatency += int(latency)
	// log.Printf("Latency: %v", latency)

	// Save log in file
	// if cacheEnable && onlyOneRequest {
	// 	SaveLog("withCache_OnlyOneRequest", fmt.Sprintf("%v", latency))
	// } else if cacheEnable && !onlyOneRequest {
	// 	SaveLog("withCache-MultipleRequest", fmt.Sprintf("%v", latency))
	// } else if !cacheEnable && onlyOneRequest {
	// 	SaveLog("noCache_OnlyOneRequest", fmt.Sprintf("%v", latency))
	// } else {
	// 	SaveLog("noCache_MultipleRequest", fmt.Sprintf("%v", latency))
	// }

	// Save log in file only when all requests are successful
	if numberOfSuccessfulRequest == numberOfRequest {
		averageLatency := totalLatency / numberOfRequest
		if cacheEnable && onlyOneRequest {
			SaveLog("withCache_OnlyOneRequest", fmt.Sprintf("Total latency (cache): %v for %v requests with average %v microsecond / request.", totalLatency, numberOfRequest, averageLatency))
		} else if cacheEnable && !onlyOneRequest {
			SaveLog("withCache-MultipleRequest", fmt.Sprintf("Total latency (cache): %v for %v requests with average %v microsecond / request.", totalLatency, numberOfRequest, averageLatency))
		} else if !cacheEnable && onlyOneRequest {
			SaveLog("noCache_OnlyOneRequest", fmt.Sprintf("Total latency (no cache): %v for %v requests with average %v microsecond / request.", totalLatency, numberOfRequest, averageLatency))
		} else {
			SaveLog("noCache_MultipleRequest", fmt.Sprintf("Total latency (no cache): %v for %v requests with average %v microsecond / request.", totalLatency, numberOfRequest, averageLatency))
		}
	}

	return nil
}

func SaveLog(name string, s ...string) {
	// open file
	path := "./result/" + name + ".txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, items := range s {
		if _, err := file.WriteString(items + "\n"); err != nil {
			panic(err)
		}
	}
}
