package sim

import (
	"fmt"
	"golang/util/io"
	"golang/util/os"
)

func ReadName() {
	fmt.Print("Hi, what's your name?: ")
	name := io.ReadString()
	fmt.Printf("Hi, %s!"+os.LineSeparator(), name)
}
