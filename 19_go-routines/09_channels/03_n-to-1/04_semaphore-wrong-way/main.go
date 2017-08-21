package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// It will block here until it can pull off a value from the 'done' channel.
	<-done
	<-done
	close(c)

	// Because the go routines we launch will not start putting value onto the channel
	// unitl we start to pull off the value from channel.(they need to synchronize.)
	//
	// To unblock above
	// we need to take value of channel c here
	// but we never get here, because we are blocked above (deadlock)
	for i := range c {
		fmt.Println(i)
	}
}
