package io

import "fmt"

func Pn(format string, args ...interface{}) {
	fmt.Printf(format+LineSeparator(), args...)
}

func P(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
