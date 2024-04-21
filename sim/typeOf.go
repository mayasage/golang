package sim

import (
	"fmt"
	"golang/util/os"
	"reflect"
)

func TypeOf() {
	order := "%-25s %v" + os.LineSeparator()

	fmt.Printf(order, "reflect.TypeOf(1): ", reflect.TypeOf(1))
	fmt.Printf(order, "reflect.TypeOf(1.1): ", reflect.TypeOf(1.1))
	fmt.Printf(order, "reflect.TypeOf('a'): ", reflect.TypeOf('a'))
	fmt.Printf(order, "reflect.TypeOf(\"a\"): ", reflect.TypeOf("a"))
	fmt.Printf(order, "reflect.TypeOf(true): ", reflect.TypeOf(true))
	fmt.Printf(order, "reflect.TypeOf(nil): ", reflect.TypeOf(nil))
}
