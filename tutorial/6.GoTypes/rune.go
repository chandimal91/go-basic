package main

import (
	"fmt"
	"reflect"
)

func main() {
	ru := 'A'
	fmt.Printf("%d \n", ru)
	fmt.Println(reflect.TypeOf(ru))
}
