// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after 

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int)  {
	x,y := 0,1
	for i:=0; i < n; i++{
		c <- x
		x,y = y, x+y
	}
	close(c)  // channel close
}

func main()  {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}