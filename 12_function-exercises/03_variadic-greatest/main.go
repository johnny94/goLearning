package main

import (
	"fmt"
)

func greatest(numbers ...int) int {
	var greatest int
	for _, v := range numbers {
		if v > greatest {
			greatest = v
		}
	}

	return greatest
}

func main() {
	x := []int{1, 2, 3}
	a := greatest(x...)

	fmt.Println(a)
}
