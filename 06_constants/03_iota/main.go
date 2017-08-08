package main

import "fmt"

const (
	_ = iota // 0, Blank indentifier (In go, _ means we don't need this).
	A = iota // 1
	B = iota // 2
	C = iota // 3
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
