package main

import (
	"fmt"
)

func main() {
	var largest int

	for i := 1; i <= 10; i++ {
		var number int
		fmt.Printf("Enter integer #%d: ", i)
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}
	}

	fmt.Printf("The largest number is: %d\n", largest)
}