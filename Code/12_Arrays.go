package main

import (
	"fmt"
)

func main() {
	// Array
	// var name_vax [number_of_elements]type
	var arr [5]int
	fmt.Println(arr)

	var arr_string [5]string
	fmt.Println(arr_string)

	// The position o the array starts to zero index.
	fmt.Println("Before: ", arr[0])
	arr[0] = 100
	fmt.Println("After: ", arr[0])

	// Different notation.
	arr_2 := [3]int{4, 5, 6}
	fmt.Println("After: ", arr_2)

	fmt.Println("After: ", len(arr_2))

	for i := 0; i < len(arr_2); i++ {
		fmt.Println(arr_2[i])
	}

	// Multidimensional arrat.
	multi := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(len(multi))

	for i := 0; i < len(multi); i++ {
		for j := 0; j < len(multi[i]); j++ {
			fmt.Println(i, j, multi[i][j])
		}
	}
}
