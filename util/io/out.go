package io

import (
	"fmt"
	"golang/util/os"
)

func P(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Pn(format string, args ...interface{}) {
	fmt.Printf(format+os.LineSeparator(), args...)
}

func Pnn(format string, args ...interface{}) {
	fmt.Printf(format+os.LineSeparator(), args...)
	fmt.Print(os.LineSeparator())
}
