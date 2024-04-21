package ui

import (
	"github.com/mbndr/figlet4go"
)

func PrettifyWord(word string) (string, error) {
	// Render maskedWord pretty
	asciiRender := figlet4go.NewAsciiRender()

	// Options
	options := figlet4go.NewRenderOptions()
	// font size
	options.FontName = "big"
	// font color
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
	}

	// Build renderStr
	return asciiRender.RenderOpts(word, options)
}
