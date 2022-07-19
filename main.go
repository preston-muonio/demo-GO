package main

import (
	"fmt"
)

var y int
var x int
var o [3]int
var p [3]int

func main() {
	var c int
	fmt.Println("Number of cells")
	fmt.Scan(&c)
	for i := 0; i < (c); i++ {
		fmt.Println("Cell:", i)

		//ask for x value
		fmt.Println("X of cell", i)
		fmt.Scan(&x)
		o[i] = x
		//ask for the y value
		fmt.Println("Y of cell", i)
		fmt.Scan(&y)
		p[i] = y
	}
	//prints grid
	for i := 0; i < 3; i++ {
		var a [3]int
		l := i
		for i := 0; i < (c); i++ {
			if l == (p[i]) {
				a[(o[i])] = 1
			}
		}
		fmt.Println(a)
	}
}
