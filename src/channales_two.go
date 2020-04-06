package main

import (
	"fmt"
)
var done chan bool = make(chan bool)

func print_greetings(source string){
	for i := 0; i<9; i++{
		fmt.Println("Hello world", i , source)
	}
	if source == "goroutine"{
		done <- true
	}
}
func main()  {
	go print_greetings("goroutine")
	print_greetings("main function")
	<- done
}

// Buffer channels are asynchronous, sending and recieving through the channel will not block
// unless the channel is full
