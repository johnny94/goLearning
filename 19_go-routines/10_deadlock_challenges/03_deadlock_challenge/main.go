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
	}()

	// Maybe this can fix the problem in 02_deadlock_challenge.
	// But we still get a error. (fatal error: all goroutines are asleep - deadlock!)
	// Why? And how to fix it?
	for {
		// Because this statement is waiting for value even the above go routine has done its works.
		fmt.Println(<-c)
	}
}
