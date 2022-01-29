package main

// several modules are imported.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	// store the input with scanner.Text().
	// strconv.ParseInt transform the input to an int number of 64 size.
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// print the typed input.
	fmt.Printf("You will be %d years old at the end of 2022", 2022-input)

}
