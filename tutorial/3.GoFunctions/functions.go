package main

import "fmt"

func main() {
	e1 := Employee{"John", "Ponting"}
	e1.fullname()
}

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullname() {
	fmt.Println(emp.fname + " " + emp.lname)
}
