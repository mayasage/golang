package sim

import (
	"fmt"
	io2 "golang/util/io"
)

func ReadName() {
	fmt.Print("Hi, what's your name?: ")
	name := io2.ReadString()
	fmt.Printf("Hi, %s!"+io2.LineSeparator(), name)
}
