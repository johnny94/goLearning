package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50 N", "99 W0", nme}
	}

	// Leave for implementation
	return 42, nil
}
