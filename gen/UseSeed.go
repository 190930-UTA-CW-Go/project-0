package gen

import (
	"math/rand"
	"time"
)

// UseSeed makes a seed to use for random numbers
func UseSeed(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	random := min + rand.Intn(max-min+1)
	return random
}
