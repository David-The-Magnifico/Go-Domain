package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	years := 5

	fmt.Printf("%-10s%-20s\n", "Interest", "Amount on Deposit")

	for rate := 0.05; rate <= 0.10; rate += 0.01 {
		fmt.Printf("%-10.2f", rate*100)

		for year := 1; year <= years; year++ {
			amount := principal * math.Pow(1.0+rate, float64(year))
			fmt.Printf("%-20.2f", amount)
		}

		fmt.Println()
	}
}