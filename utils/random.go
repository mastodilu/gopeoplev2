package utils

import (
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())
var r = rand.New(s)

// NewRandomInt returns a pseudo random int
func NewRandomInt() int {
	return rand.Int()
}

// NewRandomIntInRange returns a pseudo random int
// in the range [min, max]
func NewRandomIntInRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}
