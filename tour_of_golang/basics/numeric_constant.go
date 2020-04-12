// (An int can store at maximum a 64-bit integer, and sometimes less.) 

package main

import (
	"fmt"
)

const (
	// create a huge number by shifting a 1 bit left 100 places
	
	// In other words, the binary number that is 1 followed by 100 zero
	Big = 1 << 100
	Small = Big >> 99
)
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
