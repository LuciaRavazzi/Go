package main

import "fmt"

// To compare the value of the variables, <= >= ==

func main() {
	x := 5

	val := x < 5
	fmt.Printf("%t", val)

	val = x == 5
	fmt.Printf("\n %t", val)

	// This cannot be done.
	/* y = 4.5
	val = x == 5
	fmt.Printf("\n %t", val) */

	y := 4.5
	val = float64(x) == y
	fmt.Printf("\n %t", val)

	val = float64(x)/2 == y
	fmt.Printf("\n %t", val)

	fmt.Printf("\n %t", "tim" == "Tim")

	// the order is given by the ASCII code.
	fmt.Printf("\n %t", "a" == "b")
}
