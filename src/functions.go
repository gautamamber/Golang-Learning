package main

import "fmt"

var y = 5

func odd_even(x int)  {
	if x % 2 == 0 {
		fmt.Println("x is event")
	} else {
		fmt.Println("x is  odd")
	}
	if y % 2 == 0 {
		fmt.Println("y is event")
	} else {
		fmt.Println("y is  odd")
	}
}


func sum_of_two_numbers(x int, y int) int  {
	return x + y
}

func sum_and_difference(x int, y int) (sum int, difference int){
	return x+y, x-y
}

func multi(args ...int) int {
	sum:=0
	for i:= 0; i<len(args); i++ {
		sum+=args[1]
	}
	return sum
}

func main()  {
	var number = 4
	odd_even(number)
	sum, difference := sum_and_difference(10,4)
	fmt.Println(sum, difference)
	mul := multi(2,3,4,5)
	fmt.Println(mul)
}