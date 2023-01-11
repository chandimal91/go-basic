package main

import "fmt"

func main() {
	// local variable definition
	var x int = 10
	var y int = 20

	if x >= 10 {
		if y >= 10 {
			fmt.Printf("Inside nested if statement\n")
		}
		fmt.Printf("Value of x is : %d\n", x)
		fmt.Printf("Value of y is : %d\n", y)
	}
}
