package main

import (
	"fmt"
	"os"
)

func main() {
	arumentWithPath := os.Args  //returns all arguments including path
	arumentSlice := os.Args[1:] //returns all elements after path
	arumentAt2 := os.Args[2]    //returns specified argument only
	fmt.Println(arumentWithPath)
	fmt.Println(arumentSlice)
	fmt.Println(arumentAt2)
}
