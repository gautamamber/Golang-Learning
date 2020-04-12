// A map maps keys to values. 

// Like dictionary in python

package main

import "fmt"

type vertex struct{
	lat, long float64
}

var m map[string]vertex

func main()  {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}