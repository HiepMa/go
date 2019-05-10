package main

import "fmt"

func s1() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)
}

func s2() {
	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// func main(){
// 	s1()
// 	s2()
// }
