package main

import (
	"fmt"
)

// Data are stoed in RAM.
// The values are store in a box and the variable
// point to thre adress of that box.
// For the immutable data types, the new variable
// point to the same place of the original variable.

func main() {

	// Immutable data type:  the copy
	// operation copies the value and not the pointer.
	var x int = 5
	y := x // a copy of the value, not the reference.
	y = 7
	fmt.Println(x, y)

	// Also arrays are immutable data types.
	var my_arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Before:", my_arr)
	my_arr_copy := my_arr
	my_arr_copy[0] = -999
	fmt.Println("After:", my_arr)

	// A slice is an immutable data type: the =
	// operator bring about the same reference.
	var sl []int = []int{3, 4, 5}
	fmt.Println(sl)
	y_sl := sl
	y_sl[0] = 100
	fmt.Println(sl)
	// Also map are immutable types.
	var mp map[string]int = map[string]int{"hello": 3}
	fmt.Println(mp)
	y_mp := mp
	y_mp["hello"] = 100
	fmt.Println(mp)
}
