// A return statement without arguments returns the named return values. This is known as a "naked" return. 

package main

import (
	"fmt"
)

// (x,y int) treat as a return value

func split(sum int) (x,y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main()  {
	fmt.Println(split(18))
}