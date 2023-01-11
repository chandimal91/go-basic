package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	present := time.Now() // current time
	present.Format(time.Kitchen)
	p(present.Format("15:04:05"))
}
