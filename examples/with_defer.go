// Go has a special statement called defer that schedules a function call to be run after the function completes. 

package main

import "fmt"

func first()  {
	fmt.Println("First function")
}

func second()  {
	fmt.Println("Second Function")
}

func main()  {
	defer second()
	first()
	for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}

// A defer statement is often used with
//  paired operations like open and close, 
//  connect and disconnect, or lock and unlock 
//  to ensure that resources are released in all 
//  cases, no matter how complex the control flow
// Deferred Functions are run even if a runtime panic occurs.
// Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
