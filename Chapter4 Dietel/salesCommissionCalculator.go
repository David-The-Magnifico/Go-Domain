package main

import "fmt"

func main() {
	var numberOfItems int
	fmt.Print("Enter the number of items sold: ")
	fmt.Scanln(&numberOfItems)

	totalSales := 0.0

	for i := 1; i <= numberOfItems; i++ {
		var itemValue float64
		fmt.Printf("Enter the value of item %d: $", i)
		fmt.Scanln(&itemValue)
		totalSales += itemValue
	}

	earnings := 200 + (0.09 * totalSales)
	fmt.Printf("The salesperson's earnings for the week is: $%.2f\n", earnings)
}