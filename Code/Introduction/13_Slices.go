package main

import "fmt"

func main() {

	// Slice takes only a part of an array.
	// An array is made up of: pointer to the first element,
	// length and the capacity which are the same.
	// In this way, the next element starting from
	// the first element is reached moving forward
	// the end point.
	// A slice is a part of an array.
	// Also in this case, there is a points,
	// the length and the capacity which
	// count of many next numbers are left in the slice
	// plus the included ones.

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	var s []int = x[0:3]

	fmt.Println("Entire sequence: ", x)
	fmt.Println("Length sequence: ", len(x))
	fmt.Println("Capacity sequence: ", cap(x))

	fmt.Println("Subsequence: ", s)
	fmt.Println("Subsequence: ", len(s))
	fmt.Println("Subsequence: ", cap(s))

	// Slices in the beginning.
	var a []int = []int{5, 6, 7, 8}

	// To increase the length of the array...
	a = append(a, 10)

	// Dynamic array.
	var b []int
	fmt.Println("Dynamic array: ", b)

	c := make([]int, 5)
	fmt.Println("Make array: ", c)

}
