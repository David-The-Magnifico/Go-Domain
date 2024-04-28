package main

import "fmt"

func main() {
	citizens := []string{"John", "Mary", "Bob"}
	earnings := make([]float64, 3)
	taxes := make([]float64, 3)

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter %s's earnings for the year: $", citizens[i])
		fmt.Scanln(&earnings[i])

		if earnings[i] <= 30000 {
			taxes[i] = 0.15 * earnings[i]
		} else {
			taxes[i] = 0.15*30000 + 0.20*(earnings[i]-30000)
		}

		fmt.Printf("%s's total tax is: $%.2f\n", citizens[i], taxes[i])
	}
}