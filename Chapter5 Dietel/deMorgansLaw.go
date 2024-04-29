package main

import "fmt"

func main() {
	x, y, a, b, g, i, j := 3, 8, 5, 5, 3, 2, 7

	originalA := !(x < 5) && !(y >= 7)
	originalB := !(a == b) || !(g != 5)
	originalC := !((x <= 8) && (y > 4))
	originalD := !((i > 4) || (j <= 6))

	equivalentA := !(x < 5) || (y < 7)
	equivalentB := (a != b) && (g == 5)
	equivalentC := (x > 8) || (y <= 4)
	equivalentD := (i <= 4) && (j > 6)

	fmt.Println("Original A:", originalA, "| Equivalent A:", equivalentA)
	fmt.Println("Original B:", originalB, "| Equivalent B:", equivalentB)
	fmt.Println("Original C:", originalC, "| Equivalent C:", equivalentC)
	fmt.Println("Original D:", originalD, "| Equivalent D:", equivalentD)
}