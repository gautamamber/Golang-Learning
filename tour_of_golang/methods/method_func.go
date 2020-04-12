// Remember: a method is just a function with a receiver argument. 

package main

import (
	"fmt"
	"math"
)

type vertex struct{
	x,y float64
}	

func abs(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main()  {
	v := vertex{2,3}
	fmt.Println(abs(v))
}