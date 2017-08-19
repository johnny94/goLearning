package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()

	// Nothing happened.
	// There are three go routines in this program: main, foo and bar.
	// The main func runs foo and bar go routines then exit.
	// So, nothing will happen.
}

func foo() {
	for i := 1; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 1; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
