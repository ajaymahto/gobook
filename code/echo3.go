// There is a third way to join strings in golang
// i.e. using the strings.Join function in strings package.
package main

// gofmt tool arranges all the imports in alphabetical order
import "fmt"
import "os"
import "strings"

func main() {
	// All public function must start with a caps.
	// Incase there are two syllables we would have 
	// all separate words starting with caps.
	fmt.Println(strings.Join(os.Args[1:], " "))
}
