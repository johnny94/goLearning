package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{100, 3, 4, 1, 40, 20}
	fmt.Println(i)
	sort.Sort(sort.IntSlice(i))
	// Or we can use this instead:
	//sort.Ints(i)
	fmt.Println(i)
}
