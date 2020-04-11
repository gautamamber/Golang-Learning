package main

import (
	"fmt"
)

func main()  {
	var numbs = []int{1,2,3,4,5,6,7}
	total := 0
	for _, item := range numbs{
		total = total + item
	}
	average := total / len(numbs)
	fmt.Println("Total is", total)
	fmt.Println("Average is", average)
}
