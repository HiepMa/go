package main

import (
	"fmt"
	"./methos"
)

func main(){
	v := methos.Vertex{3, 4}
	// v.Scale(10)
	// fmt.Println(v.Abs())
	v.Scale(2)
	methos.ScaleFunc(&v, 10)

	p := methos.Vertex{3,4}
	p.Scale(3)
	methos.ScaleFunc(&p, 8)

	fmt.Println(v, p)
}