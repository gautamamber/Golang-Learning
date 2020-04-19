package main

import (
	"fmt"
)

var area int

func main()  {
	var l,b int
	fmt.Println("Enter length of rectangle")
	fmt.Scanln(&l)
	fmt.Println("Enter breadth of rectangle")
	fmt.Scanln(&b)
	area = l * b
	fmt.Println("Area of rectangle is", area)

	fmt.Println("Enter length of square")
	fmt.Scanln(&l)
	area = l * l 
	fmt.Println(area)
}