package io

import (
	"github.com/stretchr/testify/assert"
	"golang/temp"
	"golang/util/os"
	"golang/util/test"
	os2 "os"
	"path/filepath"
	"testing"
)

func TestP(t *testing.T) {
	assert.Equal(t, "Hi", test.GetStdOut(func() {
		P("Hi")
	}))
}

func TestPn(t *testing.T) {
	assert.Equal(t, "Hi"+os.LineSeparator(), test.GetStdOut(func() {
		Pn("Hi")
	}))
}

func TestPnn(t *testing.T) {
	ls := os.LineSeparator()

	assert.Equal(t, "Hi"+ls+ls, test.GetStdOut(func() {
		Pnn("Hi")
	}))
}

func TestWriteFileBytes(t *testing.T) {
	ls := os.LineSeparator()
	content := []byte("One day, there will be eternal freeze." + ls)
	dir := temp.TestDir()
	filePath1 := filepath.Join(dir, "something_just_like_this_1.txt")
	filePath2 := filepath.Join(dir, "something_just_like_this_2.txt")

	// Write with os
	err := os2.WriteFile(filePath1, content, 0666)
	if err != nil {
		panic(err)
	}

	// Write with my func
	err = WriteFileBytes(filePath2, content)
	if err != nil {
		panic(err)
	}

	// Read both files with os
	content1, err := os2.ReadFile(filePath1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, content1) // not necessary (os stuff)

	content2, err := os2.ReadFile(filePath2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, content2) // not necessary because this is os stuff)

	assert.Equal(t, content1, content2) // this is what I want to do
	// both read file content should match each other

	// Remove the created files
	err = os2.Remove(filePath1)
	if err != nil {
		panic(err)
	}

	err = os2.Remove(filePath2)
	if err != nil {
		panic(err)
	}
}

// TestWriteFileString
// Similar logic as TestWriteFileBytes
func TestWriteFileString(t *testing.T) {
	ls := os.LineSeparator()
	content := "One day, there will be eternal freeze." + ls
	dir := temp.TestDir()
	filePath1 := filepath.Join(dir, "something_just_like_this_1.txt")
	filePath2 := filepath.Join(dir, "something_just_like_this_2.txt")

	// Write with os
	err := os2.WriteFile(filePath1, []byte(content), 0666)
	if err != nil {
		panic(err)
	}

	// Write with my func
	err = WriteFileString(filePath2, content)
	if err != nil {
		panic(err)
	}

	// Read both files with os
	content1, err := os2.ReadFile(filePath1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(content1)) // not necessary (os stuff)

	content2, err := os2.ReadFile(filePath2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(content2)) // not necessary because this is os stuff)

	assert.Equal(t, string(content1), string(content2)) // this is what I want to do
	// both read file content should match each other

	// Remove the created files
	err = os2.Remove(filePath1)
	if err != nil {
		panic(err)
	}

	err = os2.Remove(filePath2)
	if err != nil {
		panic(err)
	}
}

func TestAppendFileBytes(t *testing.T) {
	ls := os.LineSeparator()
	content1 := "One day, there will be eternal freeze." + ls
	content2 := "And let there will be light" + ls
	content := content1 + content2

	dir := temp.TestDir()
	filePath1 := filepath.Join(dir, "something_just_like_this_1.txt")
	filePath2 := filepath.Join(dir, "something_just_like_this_2.txt")

	defer func() {
		// Remove the created files
		err := os2.Remove(filePath1)
		if err != nil {
			panic(err)
		}

		err = os2.Remove(filePath2)
		if err != nil {
			panic(err)
		}
	}()

	// Write both with os
	_ = os2.WriteFile(filePath1, []byte(content1), 0666)
	_ = os2.WriteFile(filePath2, []byte(content1), 0666)

	// Append with os
	file, _ := os2.OpenFile(
		filePath1,
		os2.O_WRONLY|os2.O_APPEND|os2.O_CREATE,
		0644,
	)

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err := file.WriteString(content2)

	// Append with my func
	_ = AppendFileBytes(filePath2, []byte(content2))

	// Read both files with os
	read1, _ := os2.ReadFile(filePath1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(read1)) // not necessary (os stuff)

	read2, err := os2.ReadFile(filePath2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(read2)) // not necessary because this is os stuff)

	assert.Equal(t, read1, read2) // this is what I want to do
	// both read file content should match each other
}

func TestAppendFileString(t *testing.T) {
	ls := os.LineSeparator()
	content1 := "One day, there will be eternal freeze." + ls
	content2 := "And let there will be light" + ls
	content := content1 + content2

	dir := temp.TestDir()
	filePath1 := filepath.Join(dir, "something_just_like_this_1.txt")
	filePath2 := filepath.Join(dir, "something_just_like_this_2.txt")

	defer func() {
		// Remove the created files
		err := os2.Remove(filePath1)
		if err != nil {
			panic(err)
		}

		err = os2.Remove(filePath2)
		if err != nil {
			panic(err)
		}
	}()

	// Write both with os
	_ = os2.WriteFile(filePath1, []byte(content1), 0666)
	_ = os2.WriteFile(filePath2, []byte(content1), 0666)

	// Append with os
	file, _ := os2.OpenFile(
		filePath1,
		os2.O_WRONLY|os2.O_APPEND|os2.O_CREATE,
		0644,
	)

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err := file.WriteString(content2)

	// Append with my func
	_ = AppendFileString(filePath2, content2)

	// Read both files with os
	read1, _ := os2.ReadFile(filePath1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(read1)) // not necessary (os stuff)

	read2, err := os2.ReadFile(filePath2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(read2)) // not necessary because this is os stuff)

	assert.Equal(t, read1, read2) // this is what I want to do
	// both read file content should match each other
}
