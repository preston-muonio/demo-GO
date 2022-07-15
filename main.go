package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println("read number", x, "from stdin")
	var y int
	fmt.Scan(&y)
	fmt.Println("read number", y, "from stdin")

	for i := 0; i < 3; i++ {
		var a [3]int
		if i == y {
			a[x] = 1
		}
		fmt.Println(a)
	}

}
