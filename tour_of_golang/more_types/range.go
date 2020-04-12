// The range form of the for loop iterates over a slice or map. 

package main

import "fmt"

var pow = []int{1,2,3,4,5}

func main()  {
	for i,v := range pow{
		fmt.Println(i*v)
	}

	power := make([]int, 10)

	for i := range power{
		fmt.Println(i)
	}

	for _, value := range power{
		fmt.Println(value)
	}

}