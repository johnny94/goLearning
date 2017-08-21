package main

import (
	"fmt"
)

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	plane1 := plane{}
	plane2 := plane{}
	boat1 := boat{}
	boat2 := boat{}

	rides := []vehicles{plane1, plane2, boat1, boat2}

	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
}
