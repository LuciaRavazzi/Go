package main

import "fmt"

func main() {
	// AND
	val := true && false
	fmt.Printf("%t", val)
	// OR
	val = true || false
	fmt.Printf("\n %t", val)

	// NOT
	val = true && (!false)
	fmt.Printf("\n %t", val)

	val = false || true && false
	fmt.Printf("\n %t", val)

}
