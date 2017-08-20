package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for i := range c {
			fmt.Println("A pulls off:", i)
		}
		done <- true
	}()

	go func() {
		for i := range c {
			fmt.Println("B pulls off:", i)
		}
		done <- true
	}()

	<-done
	<-done
}
