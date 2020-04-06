package main

import (
	"fmt"
)

func main()  {
	// If values not declare then values are zero
	var my_array [5]int
	fmt.Println(my_array)

	// Arrays are passed by value to functions, meaning a copy of array 
	// is passed, not a reference
	populate_integer_array(my_array)
	fmt.Println(my_array)
	my_array = populate_integer_array_return(my_array)
	fmt.Println(my_array)
	text_array()
}

func populate_integer_array(input [5]int)  {
	input[0] = 1
	input[1] = 2
	input[2] = 3
	input[3] = 4
	input[4] = 5
}

func populate_integer_array_return(input [5]int) [5]int {
	input[0] = 1
	input[1] = 2
	input[2] = 3
	input[3] = 4
	input[4] = 5
	return input
}

func text_array() {
	var input [5]string
	input[0] = "hello"
	input[1] = "world"
	input[2] = "python"
	input[3] = "c"
	input[4] = "cpp"
	fmt.Println(input)
}