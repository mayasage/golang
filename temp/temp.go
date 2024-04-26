package temp

import (
	"runtime"
)

func TestDir() string {
	switch runtime.GOOS {
	case "windows":
		return "C:\\Users\\Aithea\\OneDrive\\aithea\\notes_inprogress\\golang\\temp"
	default:
		panic("TestDir() illegal call!")
	}
}
