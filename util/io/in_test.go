package io

import (
	"github.com/stretchr/testify/assert"
	"golang/temp"
	os2 "golang/util/os"
	"golang/util/test"
	"os"
	"path/filepath"
	"testing"
)

func TestReadString(t *testing.T) {
	assert.Equal(t, "Hi", test.SetStdIn(" Hi ", func() string {
		return ReadString()
	}))
}

func TestReadInt(t *testing.T) {
	assert.Equal(t, 1, test.SetStdIn(" 1 ", func() int {
		return ReadInt()
	}))

	assert.Panics(t, func() {
		test.SetStdIn(" a ", func() int {
			return ReadInt()
		})
	})
}

func TestReadChar(t *testing.T) {
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   '\x00',
	}, test.SetStdIn("a1 ", func() ReadCharResponse {
		return ReadChar()
	}))

	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("a ", func() ReadCharResponse {
		return ReadChar()
	}))

	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   '🎈',
	}, test.SetStdIn("🎈", func() ReadCharResponse {
		return ReadChar()
	}))
}

func TestReadAlpha(t *testing.T) {
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   '\x00',
	}, test.SetStdIn("a1 ", func() ReadCharResponse {
		return ReadAlpha()
	}))
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("a ", func() ReadCharResponse {
		return ReadAlpha()
	}))
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   '🎈',
	}, test.SetStdIn("🎈", func() ReadCharResponse {
		return ReadAlpha()
	}))

	// get_lower
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_lower")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_lower")
	}))

	// get_upper
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_upper")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_upper")
	}))

	// to_upper
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'A',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("to_upper")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("to_upper")
	}))

	// to_lower
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("to_lower")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("to_lower")
	}))

	// get_lower - to_upper
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'A',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_lower", "to_upper")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_lower", "to_upper")
	}))

	// get_lower - to_lower
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_lower", "to_lower")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_lower", "to_lower")
	}))

	// get_upper - to_lower
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'a',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_upper", "to_lower")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_upper", "to_lower")
	}))

	// get_upper - to_upper
	assert.Equal(t, ReadCharResponse{
		Success: true,
		Value:   'A',
	}, test.SetStdIn("A", func() ReadCharResponse {
		return ReadAlpha("get_upper", "to_upper")
	}))
	assert.Equal(t, ReadCharResponse{
		Success: false,
		Value:   'a',
	}, test.SetStdIn("a", func() ReadCharResponse {
		return ReadAlpha("get_upper", "to_upper")
	}))
}

func TestReadFile(t *testing.T) {
	dir := temp.TestDir()
	filePath := filepath.Join(dir, "somefile.txt")

	ls := os2.LineSeparator()
	content := "Hey there Boya" + ls + "Suck this" + ls

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}

	_content, err := ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, content, _content)

	err = os.Remove(filePath)
	if err != nil {
		panic(err)
	}
}
