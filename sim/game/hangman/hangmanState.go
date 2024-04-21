package hangman

import (
	"golang/util/arr"
	"golang/util/rand"
)

type hangmanState struct {
	secretWord       string
	maskedWord       []rune
	guessLimit       int
	lastGuessed      rune
	guesses          []rune
	incorrectGuesses []rune
	prompt           func(state *hangmanState)
}

func makeRandomWord() string {
	// Make randomWord
	// Take some random words
	var randomWords = []string{
		"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
		"ZODIAC", "ZOMBIE", "FLUFF",
	}

	// Pick a random index
	randIndex := rand.Int(0, len(randomWords)-1, nil)

	// Choose the word at the random Index
	return randomWords[randIndex]
}

func makeMaskedWord(word string) []rune {
	// Make maskedWord

	// Initialize an array with '_'
	maskedWord := arr.Mk(len(word), '_')

	// Pick 2 random indices
	randIndices := arr.IntN(0, len(maskedWord)-1, 2, nil)

	// Open those indices
	for _, i := range randIndices {
		maskedWord[i] = (rune)(word[i])
	}

	return maskedWord
}

func initState() *hangmanState {
	secretWord := makeRandomWord()
	maskedWord := makeMaskedWord(secretWord)

	state := hangmanState{
		secretWord:       secretWord,
		maskedWord:       maskedWord, // Will fill in this function.
		guessLimit:       7,
		lastGuessed:      '\x00',
		guesses:          []rune{},
		incorrectGuesses: []rune{},
		prompt:           prompt,
	}

	return &state
}
