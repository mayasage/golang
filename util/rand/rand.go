package rand

import (
	"math/rand"
	"time"
)

func NewRand() *rand.Rand {
	source := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(source)
	return r
}

// Int [min, max] Open-interval, includes both.
func Int(min int, max int, src *rand.Rand) int {
	// Test Cases
	// [0,0] => 0
	// [0,1] => 0,1
	// [0,2] => 0,1,2
	// [5,10] => 5,6,7,8,9,10

	if max < min {
		panic("max must be larger than min")
	}

	if src == nil {
		src = NewRand()
	}

	// [0,0] => 0
	if min == 0 && max == 0 {
		return 0
	}

	// [0,1] => 0,1
	// [0,2] => 0,1,2
	// [5,10] => 5,6,7,8,9,10
	return src.Intn(max-min+1) + min
}
