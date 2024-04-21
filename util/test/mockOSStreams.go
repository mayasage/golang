package test

import (
	os2 "golang/util/os"
	"io"
	"os"
)

// MockStdIn
/**
 * Creates a temporary file and open Read & Write Streams to it.
 * Replaces os.Stdin with the File's ReadStream.
 * Call done() when finished.
 */
func MockStdIn() (r, w *os.File, done func()) {
	// Open Read & Write streams to Temporary File.
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	ogStdin := os.Stdin // Save original Stdin.
	os.Stdin = r        // Modify Stdin.

	return r, w, func() {
		// Done Function
		_ = r.Close()
		_ = w.Close()
		os.Stdin = ogStdin // Restore ogStdin
	}
}

// MockStdOut
/**
 * Creates a temporary file and open Read & Write Streams to it.
 * Replaces os.Stdout with the File's ReadStream.
 * Call done() when finished.
 */
func MockStdOut() (r, w *os.File, done func()) {
	// Open read & write stream to Temporary File
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	ogStdout := os.Stdout // Save original stdout
	os.Stdout = w         // Modify stdout

	return r, w, func() {
		// Done Function
		_ = r.Close()
		_ = w.Close()
		os.Stdout = ogStdout // Restore ogStdout
	}
}

// SetStdIn
// in = user input string
// cb = return cb() will be called after modifying stdin with "in"
func SetStdIn[T any](in string, cb func() T) T {
	in = in + os2.LineSeparator()

	// Mock StdIn
	_, w, done := MockStdIn()
	defer done()

	// Write to Temporary File
	_, err := w.WriteString(in)
	if err != nil {
		panic(err)
	}

	return cb()
}

// GetStdOut cb will write something to stdOut
func GetStdOut(cb func()) string {
	// Mock StdIn
	r, w, done := MockStdOut()
	defer done()

	cb()
	err := w.Close()
	if err != nil {
		panic(err)
	}

	out, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return string(out)
}
