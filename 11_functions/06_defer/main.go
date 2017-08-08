package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	// The 'world' function will not be executed until leave the function which
	// it lives in.(in this case is 'main' function)
	defer world()
	hello()
}
