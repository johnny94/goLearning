package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zero", "John", "AI", "Jenny"}
	fmt.Println(studyGroup)
	sort.StringSlice(studyGroup).Sort()
	// sort.Strings(studyGroup) <- This also works.
	fmt.Println(studyGroup)
}
