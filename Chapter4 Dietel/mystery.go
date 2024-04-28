package main

import "fmt"

func main() {
	x = -2
	total := 0
	for x <= 10 {
		y := x + 2
		x++
		total += y
		fmt.Printf("Y is: %d and total is %d\n", y, total)
	}

	
}