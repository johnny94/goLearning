package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1) // lowercase will not be exported
	fmt.Println(string(bs))
}
