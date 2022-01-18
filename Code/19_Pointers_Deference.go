package main

import "fmt"

// Pointer as the input.
func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "Change"
}

func main() {

	// * is the pointed value.
	// & is the cell reference.

	x := 7
	// Reference: id of the memory cell of the variable.
	fmt.Println("Reference: ", &x)
	fmt.Println("Value: ", x)

	// Copy the reference.
	y := &x
	fmt.Println(x, y)
	// * dereference: access to the block
	// which is pointed and modifies its value.
	*y = 8
	fmt.Println(x, y)

	// Try to change the value: only the pointer works.
	toChange := "Hello"
	changeValue2(toChange)
	fmt.Println("Without:", toChange)
	changeValue(&toChange)
	fmt.Println("With:", toChange)

	// The pointer has access to the value which is pointed.
	toChange2 := "hello"
	var pointer *string = &toChange2
	fmt.Println(pointer)

}
