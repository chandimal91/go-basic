package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("x,y,z", ","))
	fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn ", "and"))
	fmt.Printf("%q\n", strings.Split(" abc ", ""))
	fmt.Printf("%q\n", strings.Split("", "Hello"))
}
