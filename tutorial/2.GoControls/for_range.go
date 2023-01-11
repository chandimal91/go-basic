package main

import "fmt"

func main() {
	// example 01
	nums := []int{2, 3, 4}
	sum := 0

	for _, value := range nums { // "_" ignore the index
		sum += value
	}
	fmt.Println("sum:", sum)

	// example 02
	for i, num := range nums {
		if num == 2 {
			fmt.Println("index: ", i)
		}
	}

	// example 03
	kvs := map[string]string{"1": "Mango", "2": "Apple", "3": "Banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// example 04
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// example 05
	for i, c := range "RASIKA" {
		fmt.Println(i, c)
	}
}
