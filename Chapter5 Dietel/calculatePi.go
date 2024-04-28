package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	terms := 200000
	pi := 0.0
	denominator := 1.0
	found := false

	for i := 1; i <= terms; i++ {
		if i%2 != 0 {
			pi += 4.0 / denominator
		} else {
			pi -= 4.0 / denominator
		}
		denominator += 2.0

		if !found && strings.HasPrefix(strconv.FormatFloat(pi, 'f', 5, 64), "3.14159") {
			found = true
			fmt.Printf("Terms needed to reach 3.14159: %d\n", i)
		}
	}

	fmt.Printf("Value of π after %d terms: %f\n", terms, pi)
}