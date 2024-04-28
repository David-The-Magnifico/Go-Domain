package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	fmt.Println("Enter five numbers between 1 and 30:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scanln(&numbers[i])
	}

	fmt.Println("\nBar chart:")
	for _, number := range numbers {
		for j := 0; j < number; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}