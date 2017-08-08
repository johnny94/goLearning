package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"hello", "你好"}
	myOtherSlice := []string{"world", "世界", "こんにちは"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	// Remove the third element from 'mySlice'
	mySlice = append(mySlice[:3-1], mySlice[3:]...)
	fmt.Println(mySlice)
}
