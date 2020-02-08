package utils

import (
	"math/rand"
	"time"
)

// NewRandomInt returns a pseudo random int
func NewRandomInt() int {
	time.Sleep(time.Millisecond * 3)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Int()
}

// NewRandomIntInRange returns a pseudo random int
// in the range [min, max]
func NewRandomIntInRange(min, max int) int {
	time.Sleep(time.Millisecond * 3)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	out := r.Intn(max+1-min) + min
	return out
}
