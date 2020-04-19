package main

import (
	"fmt"
)

func main()  {
	var number, tempNumber, remainder int
	var result int=0
	fmt.Println("Enter three number")
	fmt.Scanln(&number)
	tempNumber = number
	for {
		remainder = tempNumber%10
		result += remainder*remainder*remainder
		tempNumber /= 10
		if tempNumber ==0{
			break
		}
	}
	if result == number{
		fmt.Println("Armstrong")
	}else{
		fmt.Println("Not armstrong")
	}
}