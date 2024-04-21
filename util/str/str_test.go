package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordCountMap(t *testing.T) {
	// Empty Word Without Options
	assert.Equal(t, map[rune]int{}, WordCountMap(""))

	// Empty Word With Options
	assert.Equal(t, map[rune]int{}, WordCountMap("", "cs"))

	// Without Options
	assert.Equal(t, map[rune]int{'H': 2, 'O': 1}, WordCountMap("Hho"))

	// With Options
	assert.Equal(
		t,
		map[rune]int{'H': 1, 'h': 1, 'o': 1},
		WordCountMap("Hho", "cs"),
	)
}

func TestCountCharInString(t *testing.T) {
	assert.Equal(t, 2, CountCharInString('a', "abca"))

	assert.Equal(
		t,
		0,
		CountCharInString('A', "abca", "case_sensitive"),
	)

	assert.Equal(
		t,
		1,
		CountCharInString('A', "abcA", "case_sensitive"),
	)
}
