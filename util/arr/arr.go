package arr

import (
	"fmt"
	"golang/util/rand"
	"math"
	rand2 "math/rand"
	"strconv"
	"strings"
)

func Contains(ints []int, int int) bool {
	for _, i := range ints {
		if i == int {
			return true
		}
	}
	return false
}

// Mk
// Use this instead of make if you want to initialize the array with a value.
func Mk[T any](size int, initVal T) []T {
	if size <= 0 {
		panic("size must be greater than zero")
	}

	res := make([]T, size)

	for i := range res {
		res[i] = initVal
	}

	return res
}

func DivideRangeIntoSlices(min int, max int, sliceSize int) []string {
	if min < 0 {
		panic("min cannot be less than 0")
	}
	if max < 0 {
		panic("max cannot be less than 0")
	}
	if max < min {
		panic("max cannot be less than min")
	}

	if sliceSize <= 0 {
		panic("limit cannot be less than 0")
	}

	// max = 0, min = 0, limit = 4, size = Ceil(max-min+1/sliceSize) = 0
	// But, the array should contain "0-0", => size = 1
	size := int(math.Ceil(float64(max-min+1) / float64(sliceSize)))

	res := make([]string, size)

	// Initially, low = min
	// low = min = 0
	low := min
	// Initially, high = (low + sliceSize - 1) = (0 + 5 - 1) = 4
	// low = 0, sliceSize = 5, high = (0 + 5 - 1) = 4
	high := math.Min(float64(low+sliceSize-1), float64(max))

	for i := 0; low <= max; i += 1 {
		res[i] = fmt.Sprintf("%v-%v", low, high)

		high += float64(sliceSize)
		if int(high) > max {
			high = float64(max)
		}

		low += sliceSize // When low will exceed max, our function ends.
	}

	return res
}

func ShuffleInPlace[T any](arr []T) {
	src := rand.NewRand()

	for i := range arr {
		randomIndex := rand.Int(0, len(arr)-1, src)
		arr[i], arr[randomIndex] = arr[randomIndex], arr[i]
	}
}

func DivideRangeIntoRandomSlices(min int, max int, sliceSize int) []string {
	result := DivideRangeIntoSlices(min, max, sliceSize)
	ShuffleInPlace(result)
	return result
}

// IntN Get "n" Int. [min, max] Open-interval, includes both.
// Always returns distinct random values.
// Panics if (max-min+1 < n).
// It is better to Panic now, that later realize "Oh Damn, it's always 0!".
func IntN(min int, max int, count int, src *rand2.Rand) []int {
	// Arguments validation
	if count <= 0 {
		panic("count must be larger than 0")
	}
	if min < 0 {
		panic("min cannot be less than 0")
	}
	if max < 0 {
		panic("max cannot be less than 0")
	}
	if max < min {
		panic("max cannot be less than min")
	}
	// Otherwise, it will cause Infinite loop.
	// Imagine, [0,0] and n = 2.
	// We wrote that we need 2 distinct integers, but there can never be 2
	// distinct integers, and thus, an Infinite loop.
	if max-min+1 < count {
		panic("[min, max] can't generate count distinct values")
	}

	// DivideRangeIntoRandomSlices
	// range of "0-10" with limit 5 will become ["0-4", "5-9", "10-10"]
	sliceSize := 5
	slicedArray := DivideRangeIntoRandomSlices(min, max, sliceSize)
	slicedArrayPointer := 0

	result := make([]int, count) // will return this
	resultPointer := 0

	// Save a randomValue used in a map.
	// Now, use this map to not use it again
	usedValues := map[int]bool{}

	// Fill result array one-by-one.
	for i := 0; i < count; i += 1 {
		for {
			// fill a random value into result array
			slice := slicedArray[slicedArrayPointer]

			lowHigh := strings.Split(slice, "-")
			low, err := strconv.Atoi(lowHigh[0])
			if err != nil {
				panic(err)
			}
			high, err := strconv.Atoi(lowHigh[1])
			if err != nil {
				panic(err)
			}

			// If it's still false after the loop, search next slice
			isExhausted := true // assume the slice is exhausted
			// try to prove it's not.
			for j := low; j <= high; j += 1 {
				if !usedValues[j] { // j is unused
					// Found a value in the slice that's not used
					isExhausted = false // This implies, the slice is not exhausted
					break
				}
			}

			// We know whether the slice contains any unused value or not.

			// Update pointer will either be used in this for loop (if the slice is
			// exhausted), or for finding the next resultPointer value.
			slicedArrayPointer = (slicedArrayPointer + 1) % len(slicedArray)

			if !isExhausted {
				// Slice contains at least 1 non-empty value.
				// So, we can fill resultPointer with an unused randomValue in [low, high]
				for {
					// Find a randomValue in [low, high]
					randomValue := rand.Int(low, high, src)

					if usedValues[randomValue] {
						// randomValue is used
						continue
						// Keep finding another random value until you find an unused one
					} else {
						// You found a random value that is unused
						usedValues[randomValue] = true
						result[resultPointer] = randomValue
						resultPointer += 1
						break
					}
				}

				break // Have found a value for resultPointer
			}
		}
	}

	return result
}
