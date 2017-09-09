package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no_file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}
