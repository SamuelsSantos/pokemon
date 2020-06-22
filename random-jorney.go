package main

import (
	"math/rand"
)

var pool = "NSEO"

func getJorney(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(3)]
	}
	return string(bytes)
}
