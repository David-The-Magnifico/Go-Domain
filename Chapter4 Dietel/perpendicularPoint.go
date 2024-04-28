package main

import {
	"fmt"
}

func main() {
	var x1, y1, x2, y2 float64

	fmt.Print("Enter the x and y coordinates of the first point: ")
	fmt.Scanln(&x1, &y1)

	fmt.Print("Enter the x and y coordinates of the second point: ")
	fmt.Scanln(&x2, &y2)

	isPerpendicular := false

	if x1 == x2 || y1 == y2 {
		isPerpendicular = true
	}

	if isPerpendicular {
		fmt.Println("The points are located on a line perpendicular to an axis.")
	} else {
		fmt.Println("The points are not located on a line perpendicular to an axis.")
	}
}