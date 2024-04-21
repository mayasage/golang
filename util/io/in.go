package io

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ReadString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ReadInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	val, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		panic(err)
	}

	return val
}

type ReadCharResponse struct {
	Success bool
	Value   rune
}

func ReadChar() ReadCharResponse {
	result := ReadCharResponse{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := scanner.Text()
	str = strings.TrimSpace(str)
	runeCount := utf8.RuneCountInString(str)

	if runeCount == 0 || runeCount > 1 {
		result.Success = false
	} else {
		result.Success = true
		result.Value, _ = utf8.DecodeRuneInString(str)
	}

	return result
}

// ReadAlpha
// options:
//   - "get_upper": Letter must be uppercase.
//   - "get_lower": Letter must be lowercase.
//   - "to_upper": Convert Letter to uppercase.
//   - "to_lower": Convert Letter to lowercase.
func ReadAlpha(options ...string) ReadCharResponse {
	// Read user input
	result := ReadChar()

	if result.Success {
		// Must be a letter
		if !unicode.IsLetter(result.Value) {
			result.Success = false
		}

		// Check if Input is uppercase/lowercase.
		for _, option := range options {
			switch option {
			case "get_upper":
				if !unicode.IsUpper(result.Value) {
					result.Success = false
					return result
				}

			case "get_lower":
				if !unicode.IsLower(result.Value) {
					result.Success = false
					return result
				}
			}
		}

		// Convert to lowercase/uppercase
		for _, option := range options {
			switch option {
			case "to_upper":
				result.Value = unicode.ToUpper(result.Value)
				break

			case "to_lower":
				result.Value = unicode.ToLower(result.Value)
				break
			}
		}
	}

	return result
}
