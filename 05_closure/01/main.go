package main

import "fmt"

func main() {
	x := 4
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Hello Scope"
		fmt.Println(y)
	}
	// fmt.Println(y) - outside scope of y
}
