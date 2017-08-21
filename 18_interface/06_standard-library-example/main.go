package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[j], a[i] = a[i], a[j] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person{
		{"Bob", 30},
		{"John", 14},
		{"Amy", 15},
		{"Judy", 10},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
