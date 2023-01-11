package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Key1"] = 10
	m["Key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)
	v, ok := m["Key2"]
	fmt.Println("The value:", v, "Present?", ok)
	i, j := m["Key4"]
	fmt.Println("The value:", i, "Present?", j)
}
