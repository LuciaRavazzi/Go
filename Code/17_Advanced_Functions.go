package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println("Hello!", x)
}

// the input is a function.
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	// Reference to the function, not calling.
	// A function can be referenced to an another variable.
	x := test
	x(5)

	// A function is referenced to another
	// variable into the main.
	my_func := func(x int) {
		fmt.Println("Hello!", x)
	}

	my_func(2)

	// In this case, the final variable is the result of the function
	// and not the function itself.
	my_func2 := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(my_func2)

	// Pass a function to a function.
	my_func3 := func(x int) int {
		return x * -2
	}
	test2(my_func3)

	// The output is a function which should be called to obtained a result.
	returnFunc("Hi")()
}
