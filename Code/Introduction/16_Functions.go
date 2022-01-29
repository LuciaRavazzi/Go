package main

import (
	"fmt"
)

// The type must be specified for both the input and the output.
func test(a, b int) (int, int) {

	if a%2 == 0 && b%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	return a + b, a - b
}

// It is possible to define the variable which contains the result.
func test2(a, b int) (z1 int, z2 int) {
	// Print something after the return operation.
	defer fmt.Println("Hello")

	z1 = a + b
	z2 = a - b

	return
}

func main() {

	var number int = 8
	add, sub := test(number, number)
	fmt.Println("Add: ", add)
	fmt.Println("Subtraction: ", sub)

	add, sub = test2(number, number)
	fmt.Println("Add: ", add)
	fmt.Println("Subtraction: ", sub)

}
