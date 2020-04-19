package main

import (
	"fmt"
)

func main()  {
	var num[100]  float64
	var temp int
	fmt.Println("Enter number of element")
	fmt.Scanln(&temp)
	for i:=0 ; i<temp;i++{
		fmt.Println("Enter number")
		fmt.Scanln(&num[i])

	}
	for j:= 1; j<temp; j++{
		if (num[0]<num[j]){
			num[0] = num[j]
		}
	}
	fmt.Println(num[0])
}