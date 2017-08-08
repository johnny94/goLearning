package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "Hello!",
		1: "Nice to see you!",
		2: "Good morning",
		3: "こんにちは",
	}

	for key, value := range myGreeting {
		fmt.Println(key, "-", value)
	}

}
