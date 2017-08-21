package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I am a regular person.")
}

func (p DoubleZero) Greeting() {
	fmt.Println("So good to see you.")
}

func main() {
	p1 := Person{
		Name: "James Bond",
		Age:  26,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "goku",
			Age:  18,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()

}
