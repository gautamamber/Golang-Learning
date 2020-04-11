package main

import (
	"fmt"
)

func main()  {
	numerator := 40
	denominator := 30

	q, r := numerator / denominator , numerator % denominator
	fmt.Println(q)
	fmt.Println(r)

}