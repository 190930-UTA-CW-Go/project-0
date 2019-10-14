package gen

import (
	"math/rand"
	"time"
)

/*
UseSeed creates a random integer between a specified minumum
and a specified maximum. It uses the current time as the seed
so it will not duplicate the same number sequence.
*/
func UseSeed(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	random := min + rand.Intn(max-min+1)
	return random
}
