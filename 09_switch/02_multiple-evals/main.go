package main

import "fmt"

func main() {
	x := "Johnny"

	switch x {
	default:
		fmt.Println("default here.")
	case "Michael", "Johnny":
		fmt.Println("Hey bro.")
	case "Andy":
		fmt.Println("Who are you?")
	}
}
