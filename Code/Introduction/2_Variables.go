package main

import "fmt"

// A variable is something which store an amount of information.
// A **static type language**, like go and C, the type of a variable cannot change.
// On the other ahdn, a **dynamic type language**, the type is allowed to be modified.

// There are different type:
// 		- uint8: 8 bit, unsigned. (0-255)
// 		- uint16: 16 bit, unsigned. (0-65535). It occupies more space.
// 		- uint32: 32 bit, unsigned.
// 		- uint64: 64 bit, unsigned.
//		- The same types can also reach negative values.
//		- Float values.
//		- Complex values.

func main() {

	//--- String.
	// type must be specified.
	var name string = "Hello Lucy"

	fmt.Println("First name: ", name)

	var name2 string
	name2 = "Lucy"

	// it could be changed.
	name2 = "Nick"

	fmt.Println("Second name: ", name2)

	// The following command gives overflow error.
	// var number uint8 = 260
	var number uint16 = 260

	number = number + 22

	fmt.Println("Hello World!", number)

}
