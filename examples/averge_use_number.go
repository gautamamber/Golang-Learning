package main

import (
	"fmt"
)

func main()  {
	var array [50]int
	var total, count int
	fmt.Println("How many number you want to print")
	fmt.Scanln(&count)
	for i :=0 ; i< count; i ++ {
		fmt.Print("Enter number")
		fmt.Scanln(&array[i])
		total += array[i]
	}
	fmt.Println(count)
	fmt.Println(total)
	average := total/count
	fmt.Println(average)
}