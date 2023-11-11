package main

import (
	"log"
	"poctest/utils"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 50},
	//{input: 1000},
	// {input: 10000},
	// {input: 382399},
}

func BenchmarkGetUserRecordV1_OneSingleUserID_CacheEnable(b *testing.B) {
	//b.Skip("skipping benchmark")
	for _, test := range table {
		threshold := 1000
		ch := make(chan bool, 1000)

		b.Run(utils.IntToString(test.input), func(b *testing.B) {
			for i := 0; i < 1000; i++ {
				go func() {
					err := GetUserRecordV1API_TDD(true, true, 9, threshold)
					if err != nil {
						log.Println(err)
						ch <- false
					} else {
						ch <- true
					}
				}()
			}

			// wait for all request to finish
			for i := 0; i < 1000; i++ {
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

func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheEnable(b *testing.B) {
	b.Skip("skipping benchmark")
	for _, test := range table {
		threshold := 10
		ch := make(chan bool, test.input)

		b.Run(utils.IntToString(test.input), func(b *testing.B) {
			for i := 0; i < test.input; i++ {
				go func() {
					err := GetUserRecordV1API_TDD(true, true, 10, threshold)
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

// func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheDisable(b *testing.B) {
// }
