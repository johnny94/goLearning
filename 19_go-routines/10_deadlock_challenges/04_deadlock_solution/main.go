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

	// Use range and don't forget to close the channel.
	for i := range c {
		fmt.Println(i)
	}
}
