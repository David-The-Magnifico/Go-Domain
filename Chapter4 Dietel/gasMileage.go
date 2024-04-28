package main

import (
	"fmt"

)

func main() {
	totalMiles := 0
	totalGallons := 0
	tripCount := 0

	var milesDriven, gallonsUsed int

	for {
		fmt.Print("Enter miles driven (or -1 to exit): ")
		fmt.Scanln(&milesDriven)

		if milesDriven == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		fmt.Scanln(&gallonsUsed)

		milesPerGallon := float64(milesDriven) / float64(gallonsUsed)
		fmt.Printf("Miles per gallon for this trip: %.2f\n", milesPerGallon)

		totalMiles += milesDriven
		totalGallons += gallonsUsed
		tripCount++

		if tripCount > 0 {
			combinedMilesPerGallon := float64(totalMiles) / float64(totalGallons)
			fmt.Printf("Combined miles per gallon: %.2f\n", combinedMilesPerGallon)
		}
	}

	fmt.Println("Exiting Gas Mileage Calculator")
}