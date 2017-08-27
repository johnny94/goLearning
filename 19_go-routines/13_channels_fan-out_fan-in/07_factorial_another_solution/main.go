package main

import (
	"fmt"
)

func main() {

	in := gen()

	// 意義上來說只要有一個 out 的管道
	// 讓所有的 go routine 都輸出到那個管道，然後直接從那個管道收集數據
	// 而不需要每個 go routine 都輸出到一個自己的管道。外面再從各個管道收集數據
	out := make(chan int)

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	fanOut(in, 10, out)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	var a string
	fmt.Scanln(&a)
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {

			out <- fact(n)
		}
	}()
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
This code was provided by my friend, Mike Van Sickle
He is an awesome programmer, an awesome Go programmer, and an awesome teacher
Check out his courses on pluralsight to learn more about Go!
*/
