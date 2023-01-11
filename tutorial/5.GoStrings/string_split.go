package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I,love,my,country"
	var arr []string = strings.Split(str, ",")
	fmt.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index : ", i, "value : ", v)
	}
}
