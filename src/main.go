package main

import (
	"./channel"
	"./rangeandclose"
	"./select"
	"fmt"
)

func main(){
	// goroutine
	//go goroutines.Say("world")
	//goroutines.Say("hello")

	//channel
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go channel.Sum(s[:len(s)/2], c)
	go channel.Sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// buffer
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//range and close
	d := make(chan int, 10)
	go rangeandclose.Fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}

	//select
	_select.Main()
	_select.Clock()
}
