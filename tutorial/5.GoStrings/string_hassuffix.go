package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "INDIA"
	fmt.Println(strings.HasSuffix(s, "DI"))
	fmt.Println(strings.HasSuffix(s, "IA"))
}
