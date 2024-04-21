package io

import (
	"github.com/stretchr/testify/assert"
	os2 "golang/util/os"
	"golang/util/test"
	"testing"
)

func TestP(t *testing.T) {
	assert.Equal(t, "Hi", test.GetStdOut(func() {
		P("Hi")
	}))
}

func TestPn(t *testing.T) {
	assert.Equal(t, "Hi"+os2.LineSeparator(), test.GetStdOut(func() {
		Pn("Hi")
	}))
}

func TestPnn(t *testing.T) {
	ls := os2.LineSeparator()

	assert.Equal(t, "Hi"+ls+ls, test.GetStdOut(func() {
		Pnn("Hi")
	}))
}
