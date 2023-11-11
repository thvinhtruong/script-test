package utils

import (
	"math/rand"
	"strconv"
)

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
