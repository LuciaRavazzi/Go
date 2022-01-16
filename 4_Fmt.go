package main

import "fmt"

// Fmt properties:
//	- %v: vaoue in default format.
//	- %T: type.
//	- %%: literal.
//	- %t: boolean.

func main() {

	// boolean type.
	fmt.Printf("Hello %t", true)

	//--- Different bases.

	// binary representation (base 2).
	fmt.Printf("\n Number: %b", 10)
	// decimal representation (base 10).
	fmt.Printf("\n Number: %d", 10)
	// eximal representation (base 16).
	fmt.Printf("\n Number: %X", 10)

	//--- Floating points.

	// Scientific notation.
	fmt.Printf("\n Number: %e", 10.347)
	// Decimal, no exponent.
	fmt.Printf("\n Number: %f", 10.347535354)
	// For large exponents: it works, the previous not.
	fmt.Printf("\n Number: %g", 10.347535354)
	// manage precision.
	fmt.Printf("\n Number: %.2f", 10.347535354)

	//--- String.

	// Not preserve the qoutes.
	fmt.Printf("\n Number: %s", "Lucy")
	// Preserve the qoutes.
	fmt.Printf("\n Number: %q", "Lucy")
}
