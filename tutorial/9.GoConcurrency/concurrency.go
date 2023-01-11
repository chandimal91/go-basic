package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	fmt.Println("start of function one")
	go fun1()
	fmt.Println("between two functions")
	go fun2()
	fmt.Println("end of function two")
	wg.Wait()
	fmt.Println("end of wait")
}

func fun1() {
	fmt.Println("inside fun1")
	for i := 0; i < 10; i++ {
		fmt.Println("fun1,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
func fun2() {
	fmt.Println("inside fun2")
	for i := 0; i < 10; i++ {
		fmt.Println("fun2,  ->", i)
		time.Sleep(time.Duration(200 * time.Millisecond))
	}
	wg.Done()
}
