package arr

import (
	"github.com/stretchr/testify/assert"
	"golang/util/rand"
	rand2 "math/rand"
	"testing"
)

func TestContains(t *testing.T) {
	assert.False(t, Contains([]int{}, 0))        // Empty Array
	assert.True(t, Contains([]int{1, 2, 3}, 1))  // Value exist in Array
	assert.False(t, Contains([]int{1, 2, 3}, 0)) // Value doesn't exist in Array
}

func TestMk(t *testing.T) {
	// panic
	assert.Panics(t, func() {
		Mk(-1, 'a')
	})

	// correct length and correct type

	// []rune length 1
	a := Mk(1, 'a')
	assert.Equal(t, 1, len(a))
	assert.Equal(t, []rune{'a'}, a)
	assert.NotEqual(t, []string{"a"}, a)

	// []string length 1
	b := Mk(1, "b")
	assert.Equal(t, 1, len(b))
	assert.Equal(t, []string{"b"}, b)
	assert.NotEqual(t, []rune{'b'}, b)

	// []int length 1
	c := Mk(1, 1)
	assert.Equal(t, 1, len(c))
	assert.Equal(t, []int{1}, c)

	// Larger size
	d := Mk(1_000_000, 1)
	assert.Equal(t, 1_000_000, len(d))
	rd := make([]int, len(d))
	for i := range rd {
		rd[i] = 1
	}
	assert.Equal(t, rd, d)
}

func TestDivideRangeIntoSlices(t *testing.T) {
	// Invalid Argument
	assert.Panics(t, func() {
		DivideRangeIntoSlices(-1, 0, 5)
	})
	assert.Panics(t, func() {
		DivideRangeIntoSlices(0, -1, 5)
	})
	assert.Panics(t, func() {
		DivideRangeIntoSlices(0, 0, 0)
	})
	assert.Panics(t, func() {
		DivideRangeIntoSlices(0, 0, -1)
	})

	// sliceSize = 1 Edge Cases
	// 0, 0, 1
	assert.Equal(
		t,
		[]string{"0-0"},
		DivideRangeIntoSlices(0, 0, 1),
	)
	// 0, 1, 1
	assert.Equal(
		t,
		[]string{"0-0", "1-1"},
		DivideRangeIntoSlices(0, 1, 1),
	)
	// 0, 2, 1
	assert.Equal(
		t,
		[]string{"0-0", "1-1", "2-2"},
		DivideRangeIntoSlices(0, 2, 1),
	)
	// 2, 2, 1
	assert.Equal(
		t,
		[]string{"2-2"},
		DivideRangeIntoSlices(2, 2, 1),
	)

	// High sliceSize doesn't matter.
	// 2, 2, 10
	assert.Equal(
		t,
		[]string{"2-2"},
		DivideRangeIntoSlices(2, 2, 10),
	)
	assert.Equal(
		t,
		[]string{"10-11"},
		DivideRangeIntoSlices(10, 11, 10),
	)
	assert.Equal(
		t,
		[]string{"10-19", "20-21"},
		DivideRangeIntoSlices(10, 21, 10),
	)

	// Random Tests
	assert.Equal(
		t,
		[]string{"0-4", "5-9", "10-10"},
		DivideRangeIntoSlices(0, 10, 5),
	)
}

func TestShuffleInPlace(t *testing.T) {
	// Check for duplicate elements
	arr := []string{"a", "b", "c", "d"}
	ShuffleInPlace(arr)
	count := map[string]int{}
	for _, v := range arr {
		count[v] += 1
	}
	for _, v := range count {
		assert.Equal(t, 1, v)
	}
}

func TestIntN(t *testing.T) {
	src := rand.NewRand()

	// [max, min, n]
	// [0, 0, 10_000]
	// (max - min + 1 < n) Panic
	assert.PanicsWithValue(
		t,
		"[min, max] can't generate count distinct values",
		func() {
			IntN(0, 0, 10_000, src)
		},
	)
	assert.PanicsWithValue(
		t,
		"[min, max] can't generate count distinct values",
		func() {
			IntN(0, 9998, 10_000, src)
		},
	)

	// n <= 0
	assert.PanicsWithValue(
		t,
		"count must be larger than 0",
		func() {
			IntN(0, 0, -1, src)
		},
	)

	// min < 0
	assert.PanicsWithValue(
		t,
		"min cannot be less than 0",
		func() {
			IntN(-1, 0, 1, src)
		},
	)

	// max < 0
	assert.PanicsWithValue(
		t,
		"max cannot be less than 0",
		func() {
			IntN(0, -1, 1, src)
		},
	)

	// [0, 9998, 10_000] Panic
	assert.Panics(t, func() {
		IntN(0, 9998, 10_000, src)
	})

	var test = func(
		min int,
		max int,
		count int,
		src *rand2.Rand,
	) ([]int, map[int]int) {
		randomValues := IntN(min, max, count, src)

		// Make a valueCount map
		valueCount := map[int]int{}
		for _, v := range randomValues {
			valueCount[v] += 1 // Increment count by 1 for each value found
		}
		assert.Equal(t, count, len(valueCount))

		// Every value in valueCount is distinct
		for _, v := range valueCount {
			assert.Equal(t, 1, v)
		}

		return randomValues, valueCount
	}

	// Step-by-Step Tests

	// min = 0
	// max = 10
	// count = 11
	_, valueCount := test(0, 10, 11, nil)
	// Every value in [0, 10] must be present in count
	for i := 0; i <= 10; i++ {
		assert.Equal(t, 1, valueCount[i])
	}

	// [0, 9999, 10_000] No Panic (9999 - 0 + 1 = 10_000)
	_, valueCount = test(0, 9999, 10_000, nil)
	// Every value in [0, 9999] must be present in count
	for i := 0; i <= 9999; i++ {
		assert.Equal(t, 1, valueCount[i])
	}

	// Bigger numbers
	// Taking 39 seconds to generate 1_000_000 distinct random values.
	// #TODO - Algorithm
	//_, valueCount = test(
	//	0,
	//	999_999,
	//	1_000_000,
	//	nil,
	//)
	//// Every value in [0, 1_000_000] must be present in count
	//for i := 0; i <= 999_999; i++ {
	//	assert.Equal(t, 1, valueCount[i])
	//}
}
