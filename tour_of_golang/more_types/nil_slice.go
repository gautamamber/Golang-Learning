// The zero value of a slice is nil. 

package main

import "fmt"

func main()  {
	var s []int

	fmt.Println(s, len(s), cap(s))
	if s == nil{
		fmt.Println("nil!")
	}
}