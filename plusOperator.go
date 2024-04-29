package main

import "fmt"

func Add(a, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return add(sum, carry)
}

func main() {
	num1 := 5
	num2 := 7
	fmt.Printf("Sum of %d and %d is: %d\n", num1, num2, add(num1, num2))
}