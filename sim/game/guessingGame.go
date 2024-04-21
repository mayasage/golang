package game

import (
	"golang/util/io"
	"golang/util/rand"
)

var pn = io.Pn

func GuessingGame() {
	num := rand.Int(0, 10, nil)
	guess := -1
	count := 0

	pn("This is a guessing game!")
	pn("Keep guessing numbers between 0 and 10 (both inclusive).")
	pn("I'll tell you how many guesses you took :)")
	pn("Let's start... Guess!")

	for guess != num {
		count += 1
		guess = io.ReadInt()

		if guess == num {
			pn("You guessed guess number!")
			pn("It took you %d tries", count)
		} else if guess > num {
			pn("Guess Lower!")
		} else {
			pn("Guess Higher!")
		}
	}
}
