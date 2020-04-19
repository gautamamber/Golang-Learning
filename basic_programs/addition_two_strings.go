package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Enter first string")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter Second string")
	var second string
	fmt.Scanln(&second)
	fmt.Println(first + second)
}