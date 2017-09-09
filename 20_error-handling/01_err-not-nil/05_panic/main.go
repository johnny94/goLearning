package main

import (
	"os"
)

func main() {
	_, err := os.Open("no_file.txt")
	if err != nil {
		// Fatalln is equivalent to Println() followed by a call to os.Exit(1)
		panic(err)
	}
}
