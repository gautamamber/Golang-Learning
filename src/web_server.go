package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "Hello world")
}

func main()  {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}