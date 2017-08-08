package main

import (
	"fmt"
)

func main() {
	var x [256]string // If we don't specify the length x will become a slice.

	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
