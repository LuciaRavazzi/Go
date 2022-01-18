package main

import "fmt"

func main() {

	// Maps ar elike distionaries in python.
	// var name map[type_key]type_vale
	var mp map[string]int = map[string]int{
		"apple":  5,
		"banana": 7,
		"orange": 10,
	}

	// In this case, the order doesn't matter.
	fmt.Println("Map:", mp)
	fmt.Println("Map length: ", len(mp))

	// Access to each element.
	fmt.Println("Map for orange:", mp["apple"])
	fmt.Println("Map for banana:", mp["banana"])
	fmt.Println("Map for orange:", mp["orange"])

	// Change values.
	mp["banana"] = 9
	fmt.Println("After Map for banana:", mp["banana"])

	// Delete values.
	delete(mp, "banaa")

	for key, item := range mp {
		fmt.Println(key, item)
	}

	// Exist statement.
	// If the key exist, return the value and true, otherwise 0 and False
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	val, ok = mp["me"]
	fmt.Println(val, ok)

}
