package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// The for-loop will loop until it found the channel is closed.
	for i := range c {
		fmt.Println(i)
	}
}
