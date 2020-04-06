package main

import (
	"fmt"
	"os"
	"log"
	"strconv" // string conversion
)

func main()  {
	var repeat int
	var err error
	if len(os.Args) >= 2 {
		repeat, err = strconv.Atoi(os.Args[1])
		if err != nil{
			log.Fatal(err)
		}
		for i:=0; i<repeat; i++{
			fmt.Println("Hello world!")
		}
	}
}