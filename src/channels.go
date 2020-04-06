// channels  are the pipe that connect cocurrent goroutines. You can send values into channels
// from one goroutines and recieve those values into another goroutines
// By default sends and receives block until both the sender and receiver are ready.
// Data flows in direction of arrow

package main

import "fmt"

func main()  {
	message := make(chan string)
	fmt.Println(message)
	go func(){
		message <- "Hey"
		fmt.Println(message)
	}()

	msg := <- message
	fmt.Println(msg)
}