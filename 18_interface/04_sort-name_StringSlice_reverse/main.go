package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zero", "John", "AI", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
