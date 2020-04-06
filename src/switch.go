package main

import "fmt"

func main(){
	z:=30
	switch z{
	case 1:
		fmt.Println("case one")
	case 2:
		fmt.Println("Case two")
	case 3:
		fmt.Println("case three")
	default:
		fmt.Println("Not equal")
	}
}