package main

import (
	"fmt"
)

func main() {
	x := 10
	changeX(&x)
	fmt.Println(x)
}
func changeX(x *int) {
	*x = *x + 1
	fmt.Println(*x)
}
