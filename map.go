package main

import (
	"fmt"
	// "golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m1 map[string]Vertex

func map1() {
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])
}

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func WordCount(s string) map[string]int {
	var st []string
	st = strings.Split(s," ")
	fmt.Println(st)
	var chuoi = make(map[string]int)
	for i:=0 ; i<len(st) ; i++ {
		chuoi[st[i]] += 1
	}
	return chuoi
}

// func main(){
// 	wc.Test(WordCount)
// }