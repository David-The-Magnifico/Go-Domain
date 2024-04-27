package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)

	if num1 == num2 {
		fmt.Println(0)
	} else if num1 > num2 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
