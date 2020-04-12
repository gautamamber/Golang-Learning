// A slice does not store any data, it just describes a section of an underlying array. 

package main

import "fmt"

func main()  {
	names := [4]string{
		"amber",
		"hello",
		"world",
		"john",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)
	b[0] = "hey man"
	fmt.Println(a,b)
	fmt.Println(names)
}