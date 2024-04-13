package io

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
