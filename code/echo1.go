// This code is an implementation of the echo command in linux
// package in Go is similar to libraries and modules
// in any other programming language.
package main

// import statements help to reuse built-in and third-party packages
// fmt - used for formatting input/output in golang
// os - used for os-level tasks
import "fmt"
import "os"


// This echo implementation simple concatenates all arguments
// that occur after the program name and then prints at the end.
// The variable sep is used to store the desired separator
// while joining the strings.
func main() {
	// Strings can also be declared like any other variable
	// the default value for the string var is a "".
	var s, sep string

	// There is only one loop in golang - for loop.
	// variants of for loop can be used to achieve the features 
	// of a while and a do-while like in any other language
	// for - for <initialization>; <condiiton>; post {}
	// os.Args is a slice which contaims the program name as the first
	// value followed by all the command-line params.
	// += is an assignment operator for the + operator
	// It is equivalent to s = s + x [s += x]
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
