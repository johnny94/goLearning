package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"` // - means not exporting this field, even it is uppercase.
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1) // lowercase will not be exported
	fmt.Println(string(bs))
}
