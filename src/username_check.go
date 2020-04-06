package main

import (
	"fmt"
	"flag"
	"log"
	"regexp"
)

const username_regex string = `^@?(\w){1,15}$`
func main()  {
	var user_input string
	flag.StringVar(&user_input, "username", "amber", "User username")
	flag.Parse()
	fmt.Println("Username validation checker")
	fmt.Println("user input", check_username_syntx(user_input))
}

func check_username_syntx(username string) bool {
	validate_result := false
	r, err := regexp.Compile(username_regex)
	if err != nil{
		log.Fatal(err)
	}

	validate_result = r.MatchString(username)
	return validate_result
}