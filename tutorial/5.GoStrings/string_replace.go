package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi...there"
	fmt.Println(strings.Replace(str, "e", "Z", 2))
}
