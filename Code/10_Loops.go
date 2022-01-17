package main

import "fmt"

func main() {

	// First loop.
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Inf condition.
	i = 0
	for true {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	// Compact notation.
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 5; i += 2 {
		fmt.Println(i)
	}

	i = 0
	for true {
		fmt.Println(i)
		i++
		if i == 10 {
			i += 2
			continue
		}

		if i == 15 {
			break
		}
	}

}
