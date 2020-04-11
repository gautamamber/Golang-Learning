package main

import (
	"fmt"
)

func main()  {
	var n1, n2 = 10, 3
	// result contain zero value by default

	result := n1 + n2
	fmt.Println("sum of two numbers", result)
	take_input_sum()
}

func take_input_sum()  {
	fmt.Println("Enter first number")
	var n1 int
	fmt.Scanln(&n1)
	fmt.Println("Enter second number")
	var n2 int
	fmt.Scanln(&n2)
	result := n1 + n2
	fmt.Println(result)
}