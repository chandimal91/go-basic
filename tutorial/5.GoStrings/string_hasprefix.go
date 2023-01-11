package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "INDIA"
	fmt.Println(strings.HasPrefix(s, "NI"))
	fmt.Println(strings.HasPrefix(s, "ND"))
	fmt.Println(strings.HasPrefix(s, "IN"))
}
