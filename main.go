package main

import (
	"log"
	"poctest/api"
	"poctest/utils"
)

// This script is to test functionality of the API of GetUserRecord v1 when applying HTTP caching and no caching
// This will test in 3 cases
// 1. GET 1 record for multiple times concurrently
// 2. GET multiple records for multiple times concurrently (small size)
// 3. GET multiple records for multiple times concurrently (large size)

const (
	httpRequestTest  = "http://localhost:9000/api/v1/GetUserRecord/1"
	testSize         = 100
	threshold        = 10
	oneRequestEnable = false
	cacheEnable      = true
)

func main() {
	testChan := make(chan bool, testSize)

	// Concurrent request
	for i := 0; i < testSize; i++ {
		go func() {
			err := GetUserRecordV1API_TDD(oneRequestEnable, cacheEnable, 1, threshold)
			if err != nil {
				testChan <- true
			} else {
				testChan <- false
			}
		}()
	}

	n := 0

	// wait for all request to finish
	for i := 0; i < testSize; i++ {
		result := <-testChan
		if result {
			log.Printf("record number %v is nil", i)
			break
		} else {
			n += 1
		}
	}

	if n == testSize {
		log.Printf("All %v requests are successfully", n)
	}
}

func GetUserRecordV1API_TDD(isOnly1Record bool, cacheEnable bool, userID int, threshold int) error {
	configuration := api.NewAPIConfig(cacheEnable)

	if isOnly1Record {
		// random 1 user id within the test size
		configuration.SetAPIEndpoint(utils.IntToString(userID))
		err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record, testSize)
		if err != nil {
			return err
		}
	} else {
		// random 10 user ids within the test size
		requestedUserId := utils.RandomInt(1, threshold)
		configuration.SetAPIEndpoint(utils.IntToString(requestedUserId))
		err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record, testSize)
		if err != nil {
			return err
		}
	}

	return nil
}
