package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(4))
}

// Challenge: Use goroutine and channel to caculate factorial
func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
