package main

import "fmt"

func main() {

	// Example 02
	fmt.Print("Enter your marks: ")
	var input int
	fmt.Scanln(&input)
	println(input)

	if input < 0 || input > 100 {
		fmt.Print("Please enter valid marks")
	} else if input >= 0 && input < 50 {
		fmt.Print("Failed")
	} else if input >= 50 && input < 60 {
		fmt.Print("D Grade")
	} else if input >= 60 && input < 70 {
		fmt.Print("C Grade")
	} else if input >= 70 && input < 80 {
		fmt.Print("B Grade")
	} else if input >= 80 && input < 90 {
		fmt.Print("A Grade")
	} else if input >= 90 && input < 100 {
		fmt.Print("A+ Grade")
	}

}
