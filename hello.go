package main

import (
	"golang/util/io"
)

func main() {
	a1 := []rune{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'}
	a2 := a1[0:5]
	io.Pn("a1: %c", a1)
	io.Pn("a2: %c", a2)
	a2[4] = 'x'
	io.Pn("a1: %c", a1)
	io.Pn("a2: %c", a2)
}
