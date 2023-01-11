package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Key1"] = 10
	m["Key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)
	fmt.Println("The value:", m["Key2"])
}
