package main

import "fmt"

func main() {
	x := 6
	y := 7
	if x > 5 && y > 5 {
		fmt.Println("x and y are > 5")
	} else {
		fmt.Println("x is <= 5")
	}
}