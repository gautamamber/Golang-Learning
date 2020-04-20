// The standard library package log provides a 
// basic infrastructure for log management in GO 
// language that can be used for logging our GO 
// programs. The main purpose of logging is to get a 
// trace of what's happening in the program, 
// where it's happening, and when it's happening. 
// Logs can be providing code tracing, profiling, 
// and analytics.

package main

import (
	"log"
)

func init()  {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Init started")
}

func main()  {
	log.Println("Main started")
	log.Fatalln("Fatal message")
	log.Panicln("Panic message")
}