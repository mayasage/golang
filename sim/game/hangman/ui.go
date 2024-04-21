package hangman

import (
	"golang/util/3p/ui"
	"golang/util/io"
	"golang/util/str"
)

var hangmanPicture = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

// playIterationIntro
// This I will play after every User Guess.
// The purpose of this is to show to the user the current state of the game,
// and prompt him/her to Guess next Letter.
func playIntro(state *hangmanState) {
	io.Pn("Let's play Hangman!")
	io.Pn("You've %d chances to find my Secret Word.", state.guessLimit)
	io.Pn("If you fail, this man's dead.")
	io.Pn("Let's start ☜(ﾟヮﾟ☜) !!!")
	io.Pn("")
}

// playFinale
// The game ends.
// Show the result.
func playFinale(state *hangmanState) {
	maskedCount := str.CountCharInString('_', (string)(state.maskedWord))

	if maskedCount > 0 {
		io.Pnn("Father will see him now ╰(*°▽°*)╯")
	} else {
		renderStr, err := ui.PrettifyWord(state.secretWord)
		if err == nil {
			io.Pn(renderStr)
		}

		io.Pnn("Here, I'll sponsor you're Bangkok trip $$$.")
	}
}

// prompt user to guess next word.
func prompt(state *hangmanState) {
	// Pick the picture based on number of incorrectGuesses
	pic := hangmanPicture[len(state.incorrectGuesses)]
	io.Pn(pic)

	// Show current state of maskedWord
	word := string(state.maskedWord)
	io.Pn("Word: %s", word)

	// If Player made any incorrectGuesses, show them
	hasIncorrectGuesses := len(state.incorrectGuesses) > 0
	if hasIncorrectGuesses {
		io.Pn("Incorrect Guesses: %s", string(state.incorrectGuesses))
	}

	io.P("Guess: ")
}

// getInput
// Prevents user from re-guessing
// This will fill "state.lastGuessed" in the end.
func getInput(state *hangmanState) {
	for {
		state.prompt(state)

		// Read user input
		input := io.ReadAlpha("to_upper")
		if input.Success == false {
			io.Pnn("You need to provide a single English Alphabet (☞ﾟヮﾟ)☞")
			continue
		}
		guess := input.Value

		// Re-ask input if Player is re-guessing.
		countInGuesses := str.CountCharInString(guess, string(state.guesses))

		if countInGuesses > 0 {
			// letter exists in guesses
			countInSecretWord := str.CountCharInString(guess, state.secretWord)

			if countInSecretWord > 0 {
				// letter exists in guesses
				// letter exists in secretWord

				if countInSecretWord > countInGuesses {
					// letter exists in guesses
					// letter exists in secretWord
					// secretWord = "aa" (2), guesses = "a" (1), guess = 'a'
					// 2 > 1, Player can guess the 2nd "a" (DONE).
					state.lastGuessed = guess
					break
				} else {
					// letter exists in guesses
					// letter exists in secretWord
					// secretWord = "a" (1), guesses = "a" (1), guess = 'a'
					// Player is re-guessing.
					io.Pnn("Guess an alphabet you haven't guessed before :>")
					continue
				}
			} else {
				// letter exists in guesses
				// letter doesn't exists in secretWord
				// Player is re-guessing.
				io.Pnn("Guess an alphabet you haven't guessed before :>")
				continue
			}
		} else {
			// letter doesn't exist in guesses (DONE).
			state.lastGuessed = guess
			break
		}
	}
}
