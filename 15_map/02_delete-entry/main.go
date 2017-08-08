package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"one":   "Hello!",
		"two":   "Nice to see you!",
		"three": "Good morning",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
