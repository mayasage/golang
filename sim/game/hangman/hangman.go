package hangman

import (
	"golang/util/io"
	"golang/util/str"
)

// useInput
func useInput(state *hangmanState) {
	guess := state.lastGuessed

	// Update guesses. Update incorrectGuesses next if Player's guess is wrong.
	state.guesses = append(state.guesses, guess)

	// Is guess correct ?
	isGuessCorrect := false
	for i, l := range state.maskedWord {
		if l == '_' {
			// find empty space in maskedWord

			// check its value in secretWord
			val := (rune)(state.secretWord[i])

			// did Player guess this ?
			if guess == val {
				// if yes, then the guess is correct.

				// fill word
				state.maskedWord[i] = guess
				isGuessCorrect = true

				io.Pnn("Bismillah Habibi... Muaah!")

				break
			}
		}
	}

	if !isGuessCorrect {
		// player guessed incorrectly
		state.incorrectGuesses = append(state.incorrectGuesses, guess)

		io.Pnn(
			"Nope. You have %d chances remaining :^",
			state.guessLimit-len(state.incorrectGuesses),
		)
	}
}

func restart(state *hangmanState) bool {
	hasGuessedWord := str.CountCharInString(
		'_',
		string(state.maskedWord),
	) == 0

	if hasGuessedWord {
		return false
	}

	hasGuessesRemaining := len(state.incorrectGuesses) < state.guessLimit

	if !hasGuessesRemaining {
		return false
	}

	return true
}

func Hangman() {
	state := initState()

	playIntro(state) // Introduce Hangman Game to Player.

	for {
		getInput(state) // Get Player's Input
		useInput(state) // Use Player's Input

		if !restart(state) { // Decides when the Game Ends.
			break
		}
	}

	playFinale(state)
}
