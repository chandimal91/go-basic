package main

import "fmt"

func main() {

	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 10:
		fmt.Print("the value is 10\n")
		fallthrough
	case 20:
		fmt.Print("the value is 20\n")
		fallthrough
	case 30:
		fmt.Print("the value is 30\n")
		fallthrough
	case 40:
		fmt.Print("the value is 40\n")
		fallthrough
	default:
		fmt.Print("it is not 10,20,30,40\n")
	}

}
