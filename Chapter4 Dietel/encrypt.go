package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter a four-digit integer: ")
	fmt.Scanln(&number)

	encrypted := encrypt(number)
	fmt.Println("Encrypted number:", encrypted)
}

func encrypt(num int) int {
	digits := make([]int, 4)

	for i := 3; i >= 0; i-- {
		digits[i] = num % 10
		digits[i] += 7
		digits[i] %= 10
		num /= 10
	}

	return (digits[2] * 1000) + (digits[3] * 100) + (digits[0] * 10) + digits[1]
}