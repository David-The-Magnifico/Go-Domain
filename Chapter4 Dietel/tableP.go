package main

import "fmt"

func main() {
	fmt.Println("N\tN^2\tN^3\tN^4")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d\t%d\t%d\t%d\n", i, i*i, i*i*i, i*i*i*i)
	}
}