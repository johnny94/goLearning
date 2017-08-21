package main

import (
	"fmt"
)

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	dog1 := dog{animal{"woof"}, true}
	cat1 := cat{animal{"meow"}, false}
	animals := []interface{}{dog1, cat1}

	fmt.Println(animals)
}
