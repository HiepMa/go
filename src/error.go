package main

import (
	"fmt"
	"math"
)

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return fmt.Sprintf(e.err)
}

func sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), run(true)
	}
	//else x<0 {
	//	re
	//}
	return math.Sqrt(-x),run(false)
}

func run(x bool) error{
	if x == true {
		return &MyError{
			"real number",
		}
	}
	return &MyError{
		err:"complex number",
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
