package main

import (
	"fmt"
)

func main() {
	// Reader
	var command string
	var number_one, number_two int
	fmt.Println("Go calculator")
	fmt.Println("Enter command")
	fmt.Println("a. Addition")
	fmt.Println("b. Substraction")
	fmt.Println("c. Multiplication")
	fmt.Println("d. Division")

	// Read command
	fmt.Scanln(&command)
	fmt.Println("You choose", command)
	// Get user number

	fmt.Println("Enter first number")
	fmt.Scanln(&number_one)
	fmt.Println("You entered", number_one)
	fmt.Println("Enter second number")
	fmt.Scanln(&number_two)
	fmt.Println("You entered", number_two)
	
	switch command {
	case "a":
		result := number_one + number_two
		fmt.Println(result)
	case "b":
		result := number_one - number_two
		fmt.Println(result)
	case "c":
		result := number_one * number_two
		fmt.Println(result)
	case "d":
		result := number_one / number_two
		fmt.Println(result)
	default:
		fmt.Println("Invalid choice")
	}

}