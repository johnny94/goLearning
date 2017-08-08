package main

import (
	"fmt"
)

func main() {
	student := make(map[string]string)
	student["name"] = "Johnny"
	student["age"] = "18"

	fmt.Println(student)

	// It is impossible to do something like this in Go.
	var a map[string]string
	a["a"] = "b"

	// This works.
	var b = make(map[string]string)
	b["b"] = "test string"
	fmt.Println(b)

	// This alse works.
	var c = map[string]string{}
	c["c"] = "c"
	fmt.Println(c)

}
