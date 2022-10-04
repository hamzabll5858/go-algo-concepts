package main

import (
	"math/rand"
	"time"
)

func GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
