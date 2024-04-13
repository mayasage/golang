package sim

import (
	"golang/util/io"
)

var pn = io.Pn

func FormatOut() {
	// ----- FORMATTED PRINT -----
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	pn("%s %d %c %f %t %o %x", "Stuff", 1, 'A', 3.14, true, 1, 1)

	// Float formatting
	pn("%9f", 3.14)      // Width 9
	pn("%.2f", 3.141592) // Decimal precision 2
	pn("%9.f", 3.141592) // Width 9 no precision
}
