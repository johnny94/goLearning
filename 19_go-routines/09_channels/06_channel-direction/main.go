package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	cSum := puller(c)
	for i := range cSum {
		fmt.Println(i)
	}
}

// <-chan means we can only recive value from the channel
// ref: https://golang.org/ref/spec#Channel_types
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
