package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi
	var x = "10"
	var y = 20
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)

	// Itoa
	a := 10
	b := "I have this many: " + strconv.Itoa(a)
	fmt.Println(b)
}
