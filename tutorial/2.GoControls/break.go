package main

import "fmt"

func main() {
	// example 01
	var a int = 1
	for a < 10 {
		fmt.Print("Value of a is ", a, "\n")
		a++
		if a > 5 {
			break
		}
	}

	// example 02
	var x int
	var y int
	for x = 1; x < 4; x++ {
		for y = 1; y < 4; y++ {
			if x == 2 && y == 2 {
				break
			}
			fmt.Print(x, " ", y, "\n")
		}
	}
}
