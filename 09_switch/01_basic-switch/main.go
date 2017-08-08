package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("Hey guys.")
		break // break is not necessary in Go.
	case 2:
		fmt.Println("Greetings!")
		fallthrough // After seeing 'fallthrough', case 3 will also be executed.
	case 3:
		fmt.Println("What do you want?")
	default: // default is also not necessary in Go.
		fmt.Println("eh, What?")
	}
}
