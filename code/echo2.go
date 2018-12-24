// In this implementation of the echo command we will 
// use range to iterate over the slice of parameters.
package main

import "fmt"
import "os"

func main() {
	// Shorthand operators can be used to perform 
	// initialization and declaration together.
	// This invokes the compiler to infer the type
	// of the variable used.
	s, sep := "", ""

	// range function iterates over the iterable and returns
	// two values [index, value] and [key, value] in case of map.
	// The _ variable or the empty variable holds values that are
	// ignored as go does not allow unused variables in other words 
	// clubs all useless values into one variable.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
