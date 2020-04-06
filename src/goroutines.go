// A goroutine is a lightweight thread of execution. Runs a program in concurrent manner
// When program is terminated, goroutines also terminated

package main

import (
	"fmt"
	"time"
)

func f(from string)  {
	for i:=0; i<4; i++ {
		fmt.Println(from, ":", i)
	}
}

func main()  {
	f("Hello")
	// Go routine
	go f("goroutine")

	go func (msg string)  {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}