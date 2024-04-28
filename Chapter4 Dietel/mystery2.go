package main

import "fmt"

func main() {
	count := 1
	for count <= 20 {
		if count%3 == 1 {
			fmt.Println("########")
		} else {
			fmt.Println("++++++++")
		}
		count++
	}
}