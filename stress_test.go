package main

import (
	"log"
	"poctest/api"
	"poctest/utils"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	// {input: 382399},
}

func BenchmarkGetUserRecordV1_OneSingleUserID_CacheEnable(b *testing.B) {
	for _, test := range table {
		threshold := 1000
		ch := make(chan bool, test.input)

		b.Run(utils.IntToString(test.input), func(b *testing.B) {
			for i := 0; i < test.input; i++ {
				go func() {
					err := GetUserRecordV1API_TDD(true, false, 2, threshold)
					if err != nil {
						ch <- false
					} else {
						ch <- true
					}
				}()
			}

			// wait for all request to finish
			for i := 0; i < test.input; i++ {
				result := <-ch
				if !result {
					log.Printf("record number %v is nil", i)
					b.Fail()
					break
				}
			}
		})
	}
}

// func BenchmarkGetUserRecordV1_OneSingleUserID_CacheDisable(b *testing.B) {
// }

// func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheEnable(b *testing.B) {
// }

// func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheDisable(b *testing.B) {
// }

func GetUserRecordV1API_TDD(isOnly1Record bool, cacheEnable bool, userID int, threshold int) error {
	configuration := api.NewAPIConfig(cacheEnable)

	if isOnly1Record {
		// random 1 user id within the test size
		configuration.SetAPIEndpoint(utils.IntToString(userID))
		err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record)
		if err != nil {
			return err
		}
	} else {
		// random 10 user ids within the test size
		requestedUserId := utils.RandomInt(1, threshold)
		configuration.SetAPIEndpoint(utils.IntToString(requestedUserId))
		err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record)
		if err != nil {
			return err
		}
	}

	return nil
}
