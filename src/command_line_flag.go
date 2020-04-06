package main

import (
	"fmt"
	"flag"
)

func main()  {
	var name string
	flag.StringVar(&name, "amber", "hello", "The name is amber")
	flag.Parse()
	fmt.Println(name)
}