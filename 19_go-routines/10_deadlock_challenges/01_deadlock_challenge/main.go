package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Block until someone start to receive value
	c <- 1

	// Receive value from channel.
	// But the program never reaches here, because it is blocked above. Deadlock
	fmt.Println(<-c)
}
