package main

import "fmt"

func main() {
	a := 5

	fmt.Println(a)
	fmt.Println(&a)

	var b *int
	b = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(a)

}
