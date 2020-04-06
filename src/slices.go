// Slice is a reference type

package main

import (
	"fmt"
)

func main()  {
	// Here we use make function to create new slice of length 5.
	var my_slice []int = make([]int, 5)
	fmt.Println("content os my slice", my_slice)
	populate_integer_slice(my_slice)
	fmt.Println(my_slice)
	// len() to check length of slice
	fmt.Println(len(my_slice))
	// cap() capacity of slice
	fmt.Println("capacity", cap(my_slice))
	fmt.Println("Add new element in slice")
	my_slice = append(my_slice, 18)
	fmt.Println(my_slice)
	fmt.Println(len(my_slice))
	// cap() capacity of slice
	fmt.Println("capacity", cap(my_slice))
	// slicing using shorthand notation
	s := my_slice[1:4]
	fmt.Println(s)

	// when you change in sub slice, change will reflect in original one
	s[1] = 7
	fmt.Println(my_slice)
	change_line_example()

}

func populate_integer_slice(input []int)  {
	input[0] = 0
	input[1] = 1
	input[2] = 2
	input[3] = 3
	input[4] = 4

}


func change_line_example()  {
	fmt.Print("Original lineup")
	inx := []string{"hello", "world", "jon", "tim"}
	fmt.Println(inx, "\n")

	inx = append(inx[:0], inx[1:]...)
	fmt.Println(inx)
}