package main

import "fmt"

func main()  {
	sum := 1
	for ; sum < 100; {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}