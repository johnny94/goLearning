package main

import (
	"fmt"
)

func main() {
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.0"
	student1[3] = "75.0"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Johnny"
	student2[1] = "Wang"
	student2[2] = "90.0"
	student2[3] = "80.0"

	records = append(records, student2)

	fmt.Println(records)
}
