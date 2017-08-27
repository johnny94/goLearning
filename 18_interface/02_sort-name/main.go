package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[j], p[i] = p[i], p[j] }

func main() {
	studyGroup := people{"Zero", "John", "AI", "Jenny"}
	fmt.Println(len(studyGroup))
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
