package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no_file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
}
