package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := person{age: 30, firstName: "Rasika", lastName: "Chandimal"}
	fmt.Println(x)
	fmt.Println(x.firstName)
	fmt.Println(x.lastName)
	fmt.Println(x.age)
}
