package io

import (
	"fmt"
	os2 "golang/util/os"
	"os"
)

func P(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Pn(format string, args ...interface{}) {
	fmt.Printf(format+os2.LineSeparator(), args...)
}

func Pnn(format string, args ...interface{}) {
	fmt.Printf(format+os2.LineSeparator(), args...)
	fmt.Print(os2.LineSeparator())
}

// WriteFileBytes
// Creates new file if it doesn't exist. Overwrite otherwise.
func WriteFileBytes(filePath string, content []byte) error {
	err := os.WriteFile(filePath, content, 0666)
	return err
}

// WriteFileString
// Creates new file if it doesn't exist. Overwrite otherwise.
func WriteFileString(filePath string, content string) error {
	bytes := []byte(content)
	err := os.WriteFile(filePath, bytes, 0666)
	return err
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// AppendFileBytes
// Create new file if it doesn't exist. Append otherwise.
// Options: os.O_WRONLY|os.O_APPEND|os.O_CREATE
// - os.O_WRONLY: write-only
// - os.O_APPEND: append instead of overwriting
// - os.O_CREATE: create file if it doesn't exist
func AppendFileBytes(filePath string, content []byte) error {
	file, err := openFile(filePath)

	defer func() {
		_err := file.Close()
		if _err != nil {
			err = _err
		}
	}()

	// Append data to the file
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return err
}

// AppendFileString
// Create new file if it doesn't exist. Append otherwise.
// Options: os.O_WRONLY|os.O_APPEND|os.O_CREATE
// - os.O_WRONLY: write-only
// - os.O_APPEND: append instead of overwriting
// - os.O_CREATE: create file if it doesn't exist
func AppendFileString(filePath string, content string) error {
	file, err := openFile(filePath)

	defer func() {
		_err := file.Close()
		if _err != nil {
			err = _err
		}
	}()

	// Append data to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return err
}
