// Go was born in order to run fast as C and has the simplicity of python.

// Go is a compiled language which means that before the execution, it should be
// transformed in something that the computer can manage, like the assembly language.
// go run is the command which compile the script
// and automatically run the file. These steps could be splitted with go build
// and then, the result is runned with a different command.
// go.exe is the build or compiled file which can be runned (even without go compiler!).

// Go needs packages.
package main

// fmt is a standard module which print things out to the console.
import "fmt"

// Entry point to our program.
// This function is automatically run when the program is runned.
func main() {
	fmt.Println("Hello World!")
}
