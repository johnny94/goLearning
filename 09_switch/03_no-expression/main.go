package main

import "fmt"

func main() {
	name := "Johnny"

	switch {
	case len(name) == 2:
		fmt.Println("Hey bro.")
	case name == "Andy":
		fmt.Println("Who are you?")
	default:
		fmt.Println("No matched; This is the default.")
	}
}
