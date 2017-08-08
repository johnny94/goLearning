package main

import "fmt"

func main() {

	if err := doSomething(); err {
		fmt.Println("Something wrong happened")
	}
}

func doSomething() bool {
	return true
}
