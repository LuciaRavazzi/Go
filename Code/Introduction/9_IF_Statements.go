package main

import "fmt"

func main() {

	name := "Nick"

	// If statement.

	fmt.Println("Before if")

	if name == "Mine" {

		fmt.Println("My name is Lucy. Hi :)")

	} else {

		fmt.Println("My name is not Lucy. !Hi :)")
	}

	fmt.Println("After if")

	// Nested If statement.

	age := 90

	if age < 17 {
		fmt.Println("You can watch television.")
	} else if age < 14 {
		fmt.Println("You cannot posses a phone.")
	} else {
		fmt.Println("Maybe you're an adult.")
	}

}
