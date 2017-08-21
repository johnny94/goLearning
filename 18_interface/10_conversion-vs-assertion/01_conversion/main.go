package main

import (
	"fmt"
)

func main() {
	var x = 12
	var y = 12.134687
	fmt.Println(float64(x) + y) // int to float64
	fmt.Println(x + int(y))     // float64 to int

	var a rune = 'a' // rune is an alias for int32
	var b int32 = 'b'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(string(a))
	fmt.Println(string(b))

	// conversion: []byte to string
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))

	// conversion: string to []byte
	fmt.Println([]byte("hello"))
}
