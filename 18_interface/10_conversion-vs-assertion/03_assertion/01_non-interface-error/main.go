package main

import (
	"fmt"
)

func main() {
	// name := "Sydeny" This doesn't work.
	var name interface{} = "Sydeny"

	// Check if the concrete type of this 'interface{}' is "string'.
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
