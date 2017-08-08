package main

import "fmt"

func main() {
	answer := (true && false) || (false && true) || !(false && false)
	fmt.Println(answer)
}
