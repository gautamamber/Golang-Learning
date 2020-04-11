// func Replace(s, old, new string, n int) string  

package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "This is test example"
	fmt.Println("Original string", str)
	fmt.Println("Output is", strings.Replace(str, " ", "", -1))
	fmt.Println("Using TrimSpace", strings.TrimSpace(str))
}
