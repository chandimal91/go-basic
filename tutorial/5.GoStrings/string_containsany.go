package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "h & r"))
	fmt.Println(strings.ContainsAny("Hello", "er"))
	fmt.Println(strings.ContainsAny("", ""))
}
