// A slice literal is like an array literal without the length. 

package main

import "fmt"

func main()  {
	q := []int{1,2,3,4,5,6}
	fmt.Println(q)

	e := []bool {true, false, true, true, false, true}
	fmt.Println(e)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}