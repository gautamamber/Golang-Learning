package main

import (
	"fmt"
)

func main()  {
	var country map[int] string
	fmt.Println(country)
	country = make(map[int] string)
	fmt.Println(country)
	country[1] = "India"
	country[2] = "USA"
	country[3] = "UK"
	country[4] = "Aus"
	
	for i,j := range country{
		fmt.Println(i,j)
	}
}