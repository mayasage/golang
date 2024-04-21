package rand

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	src := NewRand()

	// [0, 0]
	for i := 0; i < 10_000; i++ {
		r := Int(0, 0, src)
		assert.Equal(t, 0, r) // 0 every time
	}

	// min > max
	for i := 0; i < 10_000; i++ {
		assert.Panics(t, func() { Int(1, 0, src) })
	}

	// [0, 1]
	for i := 0; i < 10_000; i++ {
		r := Int(0, 1, src)
		assert.Contains(t, []int{0, 1}, r) // 0 or 1
	}
}
