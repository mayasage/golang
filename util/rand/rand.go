package rand

import (
	"math/rand"
	"time"
)

func Int(min int, max int) int {
	source := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(source)
	return min + r.Intn(max-min)
}
