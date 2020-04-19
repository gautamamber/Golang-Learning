// A goroutine is a lightweight thread managed by the Go runtime. 

// Thread in Golang

// example: go f(x, y, z)

package main

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i<5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	
}

func main()  {
	go say("hello")
	say("world")
}