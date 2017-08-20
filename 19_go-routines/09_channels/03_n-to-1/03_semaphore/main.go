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

	go func() {
		<-done // It will pause here until it can pull off a value from the 'done' channel.
		<-done
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
