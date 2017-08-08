package main

import "fmt"

func main() {

	// Zero-value
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Println()

	// short hand (can only be used inside func)
	e := 123
	f := "Hello"
	g := 4.17
	h := true

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

}
