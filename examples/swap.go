package main

import (
	"fmt"
)

func main()  {
	num1, num2 := 12,23
	fmt.Println("Before swap")
	fmt.Println("Number one", num1)
	fmt.Println("Number two", num2)

	// Swap

	temp := num1
	num1 = num2
	num2 = temp

	fmt.Println("After swap")
	fmt.Println("Number one", num1)
	fmt.Println("Number two", num2)
	// Another method

	num1, num2 = num2, num1  

	fmt.Println("Again swap")
	fmt.Println("Number one", num1)
	fmt.Println("Number two", num2)

}