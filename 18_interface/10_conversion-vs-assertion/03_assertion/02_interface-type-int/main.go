package main

import (
	"fmt"
)

func main() {
	var i interface{} = 10
	fmt.Printf("%T\n", i)

	// Even the last statement says the type of i is 'int',
	// we still cannot use it to plus a 'int' directly.
	fmt.Println(6 + i)       // This doesn't work.
	fmt.Println(6 + i.(int)) // This works.
}
