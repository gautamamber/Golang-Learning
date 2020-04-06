package main

import "fmt"

func main(){
	// Integer without value is zero
	var my_integer int32
	fmt.Println(my_integer)

	var my_float float64 = 36.333
	fmt.Println(my_float)

	x,y,z := 2,3,4
	fmt.Println(x,y,z)

	complex_number := 5 + 2i
	fmt.Println(complex_number)
	
	const (
		speed_of_light = 233333
		pi = 3.14
		fav_number = 108
	)

	var (
		a int = 0
		b = 22
		c = 3.3
	)
	fmt.Println(a,b,c,speed_of_light, pi, fav_number)

	var string_value string = "amber"
	fmt.Println(string_value)
	fmt.Println(len(string_value))
}