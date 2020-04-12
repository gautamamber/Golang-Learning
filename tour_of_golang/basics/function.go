package main

import "fmt"

func main()  {
	fmt.Println(add(23,33))
	fmt.Println(add_cont(33,22))
}
func add(a int, b int) int {
	result := a+b
	return result
}

func add_cont(x int,y int) int  {
	result := x+y
	return result
}