// Rating system
// Take  rating from user and store with time stamp

package main

import (
	"fmt"
	"bufio" // take input from user
	"os"
	"strconv" // convert str to number
	"strings"
	"time"
)

func main()  {
	var name string
	var userRating string

	// Front end (Dummy)
	// Take name as input

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name...")
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to int or fload
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our service between 1 and 5... ")
	userRating, _ = reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)


	// Backend

	fmt.Printf("Hello %v, Thanks for rating out Service with %v start rating", name, mynumrating, time.Now().Format(time.Stamp))

	if mynumrating == 5 {
		fmt.Println("Bonus for team 5 start service")
	} else if mynumrating == 4 || mynumrating == 3 {
		fmt.Println("we are always improving")
	} else {
		fmt.Println("Need serious work from our side")
	}

}