package main

import "fmt"

func main() {

	ans := -1

	// Switch.
	switch ans {
	case 1, -1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Not a case")
	}

	// Switch.
	switch {
	case ans > 0:
		fmt.Println("Greater than zero")
	case ans == 0:
		fmt.Println("Equal to zero")
	case ans < 0:
		fmt.Println("Less than zero")

	}

}
