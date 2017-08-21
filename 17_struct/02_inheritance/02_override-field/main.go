package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   18,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	fmt.Println(p1.First)
	fmt.Println(p1.Person.First)
}
