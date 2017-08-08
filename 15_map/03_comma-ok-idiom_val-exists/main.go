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

	//delete(myGreeting, 2)

	if val, exist := myGreeting[2]; exist {
		fmt.Println("That value exists.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exist)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exist)
	}

	fmt.Println(myGreeting)

}
