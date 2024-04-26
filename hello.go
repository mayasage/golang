package main

import (
	"golang/util/io"
	"os"
)

func main() {
	io.Pn("%v", os.Args[1:])
}
