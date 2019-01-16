package main

import "fmt"

func main() {
	fmt.Println(fmt.Sprintf("%s", "test"))
	a()
}

func a() int {
	var a1 int
	a1 = 2
	return a1
}
