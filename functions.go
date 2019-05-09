package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	sum :=0
	return func(i int) int {
		if i < 0 {
			return 0
		}
		switch i {
			case 0:{
				return 0
			}
			case 1: {
				return 1
			} 
			default:{
				f :=fibonacci()
				sum = f(i-1) + f(i-2)
				return sum
			}
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

