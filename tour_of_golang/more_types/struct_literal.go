// A struct literal denotes a newly allocated struct value by listing the values of its fields. 

package main

import "fmt"

type vertex struct {
	x , y int
}

var (
	v1 = vertex{1,2}
	v2 = vertex{322,3}
	v3 = vertex{x:1}
	v4 = vertex{}
	p = &vertex{1,3}
)
func main()  {
	fmt.Println(v1, v2, v3, v4, p)
}