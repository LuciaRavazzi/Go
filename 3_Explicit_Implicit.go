package main

import "fmt"

func main() {

	// Explicit declaration.
	// var number uint16 = 260

	// Implicit declaration: go infer the type.
	var number2 = 260
	// it is generic.
	fmt.Printf("\n %T", number2)

	var number3 = 260.1
	fmt.Printf("\n %T", number3)

	// omit the var word and infer the time.
	number4 := 6
	fmt.Printf("\n %T", number4)
	// number4 := "hello" doesn't work

	boolean := false
	fmt.Printf("\n %T", boolean)

	// a new variable int is initialize to zero.
	var new_number uint16
	fmt.Printf("\n %T", new_number)

}
