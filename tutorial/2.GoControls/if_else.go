package main

import "fmt"

func main() {
	// Example 01
	var a int = 11
	if a%2 == 0 {
		fmt.Printf("a is even\n")
	} else {
		fmt.Printf("a is odd\n")
	}
	fmt.Printf("value of a is : %d\n", a)

	// Example 02
	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)

	if input%2 == 0 {
		fmt.Printf(" is even\n")
	} else {
		fmt.Printf(" is odd\n")
	}
	fmt.Printf("value of input is : %d\n", input)

}
