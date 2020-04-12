// You can declare a method on non-struct types, too. 

package main

import (
	"fmt"
	"math"
)

type myfloat float64

func (f myfloat) abs() float64 {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

func main()  {
	f := myfloat(-math.Sqrt2)
	fmt.Println(f.abs())
}