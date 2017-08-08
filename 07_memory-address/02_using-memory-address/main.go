package main

import "fmt"

const Pi = 3.14

func main() {
	var radius float64
	fmt.Printf("Enter the radius: ")
	fmt.Scan(&radius)
	area := radius * radius * Pi
	fmt.Println("The area of the circle which radius is ", radius, " is ", area)
}
