package api

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func MakeCallToApi(endpoint string, cacheEnable bool, onlyOneRequest bool) error {
	startTime := time.Now()
	_, err := http.Get(endpoint)
	if err != nil {
		log.Printf("Error when making call to API: %s", err)
		return err
	}

	latency := time.Since(startTime)

	//log.Printf("Latency: %s", latency)

	go func() {
		// Save log in file
		if cacheEnable && onlyOneRequest {
			saveLog("withCache_OnlyOneRequest", latency.String())
		} else if cacheEnable && !onlyOneRequest {
			saveLog("withCache-MultipleRequest", latency.String())
		} else if !cacheEnable && onlyOneRequest {
			saveLog("noCache_OnlyOneRequest", latency.String())
		} else {
			saveLog("noCache_MultipleRequest", latency.String())
		}
	}()

	return nil
}

func saveLog(name string, s ...string) {
	// open file
	var mutex sync.Mutex
	mutex.Lock()
	path := "./log/" + name + ".log"
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
	mutex.Unlock()
}
