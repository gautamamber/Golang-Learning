package main

import "fmt"

const (
	_ = iota
	// _ is a blank indentifier
	// It will generate series of 1, 2,3 
	red_light
	gree_light
	yello_light
)

func main()  {
	fmt.Println("red light", red_light)
	fmt.Println("green light", gree_light)
	fmt.Println("yellow light", yello_light)
}