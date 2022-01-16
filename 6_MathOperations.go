package main

import (
	"fmt"
)

func main() {

	// Integers.
	var num1 int = 9
	var num2 int = 4

	answer := num1 / num2
	// ruìound result.
	fmt.Println("Ratio between int: ", answer)

	// Int e float.
	var num3 float32 = 8.1
	var num4 int = 4
	// Invalid operation! The type must be the same.
	//answer = num3 / num4
	// fmt.Println("Ratio between int: ", answer)
	fmt.Println("Ratio between int and float: ", num3/float32(num4))

	fmt.Println("Mod result: ", int16(num3)%int16(num4))

}
