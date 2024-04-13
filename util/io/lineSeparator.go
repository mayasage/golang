package io

import "runtime"

func LineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	case "linux":
		return "\n"
	case "darwin":
		return "\r"
	default:
		return "\n"
	}
}
