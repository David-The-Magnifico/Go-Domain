package main

import "fmt"

func main() {
	var currentPopulation, growthRate float64

	fmt.Print("Enter the current world population: ")
	fmt.Scanln(&currentPopulation)

	fmt.Print("Enter the annual world population growth rate (in percentage): ")
	fmt.Scanln(&growthRate)

	population := currentPopulation
	year := 0

	for population < 2*currentPopulation {
		population = population * (1 + growthRate/100)
		year++
	}

	fmt.Printf("The world population will double in %d years.\n", year)
}