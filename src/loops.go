package main

import (
	"fmt"
	"time"
)
func main()  {
	// classic loop
	for i:=0; i<10;i++ {
		if i == 0{
			continue
		}
		fmt.Println(i)
	}
	
	fmt.Println("\n\n")

	// signle condition for loop
	j := -20
	for j!=0{
		fmt.Println("single condition", j)
		j++
	}
	fmt.Println("\n\n")

	loop_timer := time.NewTimer(time.Second * 9)
	// An infinite loop, will terminate after 9 second
	for {
		fmt.Println("Inside the infinite loop")
		// Listen the timer
		<-loop_timer.C
		break
	}
}