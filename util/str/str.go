package str

import "unicode"

// Supported Options
type wordCountOptions struct {
	caseSensitive bool
}

func parseOptionsString(options string) wordCountOptions {
	// default options
	parsedOptions := wordCountOptions{
		caseSensitive: false,
	}

	// modify default options
	if options == "cs" {
		parsedOptions.caseSensitive = true // true
	}

	// return parsedOptions
	return parsedOptions
}

// WordCountMap
// options is optional. This is why I wrote ...string.
// In reality, it will only use the first string.
// The first string will be treated as options.
// Send only 1 string as options if you want.
func WordCountMap(word string, options ...string) map[rune]int {
	res := map[rune]int{}

	// Parse Options
	var parsedOptions wordCountOptions
	if len(options) > 0 {
		parsedOptions = parseOptionsString(options[0])
	} else {
		parsedOptions = parseOptionsString("")
	}

	for _, c := range word {
		// If caseSensitivity is False, then 'a' & 'A' are same.
		// If caseSensitivity is True, then 'a' & 'A' are different.
		if !parsedOptions.caseSensitive {
			c = unicode.ToUpper(c)
		}

		res[c] += 1
	}

	return res
}

func CountCharInString(char rune, str string, options ...string) int {
	count := 0

	isCaseSensitive := false
	for _, o := range options {
		switch o {
		case "case_sensitive":
			isCaseSensitive = true
		}
	}

	// If case is insensitive, then make both char and c in str upper case.
	if !isCaseSensitive {
		char = unicode.ToUpper(char)
	}
	for _, c := range str {
		if !isCaseSensitive {
			c = unicode.ToUpper(c)
		}
		if c == char {
			count += 1
		}
	}

	return count
}
