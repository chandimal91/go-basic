package main

import "fmt"

func main() {
	var input int
Lable:
	fmt.Print("You are not eligible to vote\n")
	fmt.Print("Please enter your age: ")
	fmt.Scanln(&input)
	if input <= 17 {
		goto Lable
	} else {
		fmt.Print("You can vote")
	}

}
