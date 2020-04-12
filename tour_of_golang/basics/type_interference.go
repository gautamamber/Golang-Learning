// When declaring a variable without 
// specifying an explicit type 
// (either by using the := syntax 
// or var = expression syntax), 
// the variable's type is inferred 
// from the value on the right hand 
// side. 

// var i int
// j := i // j is an int

// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128

package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
