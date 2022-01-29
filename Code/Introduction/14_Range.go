package main

import "fmt"

func main() {

	var a []int = []int{1, 2, 3, 4}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, elem := range a {
		fmt.Println(i, elem)
	}
}
