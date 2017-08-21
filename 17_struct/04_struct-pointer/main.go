package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.age)  // equivalent to (*p1).age in go
	fmt.Println(p1.name) // equivalent to (*p1).name in go
}
