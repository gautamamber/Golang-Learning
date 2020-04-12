// Go does not have classes. 

// However, you can define methods on types. 

// A method is a function with a special receiver argument. 

// The receiver appears in its own argument list between the func keyword and the method name. 



package main

import (
	"fmt"
	"math"
)

type vertex struct{
	x,y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	v := vertex{2,4}
	fmt.Println(v.abs())
}

