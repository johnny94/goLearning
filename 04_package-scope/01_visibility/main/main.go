package main

import (
	"fmt"

	"github.com/johnny94/goLearning/04_package-scope/01_visibility/vis"
)

func main() {
	vis.PrintVar()
	fmt.Println(vis.MyName)
}
