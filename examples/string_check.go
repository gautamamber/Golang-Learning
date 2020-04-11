package main

import (
	"fmt"
)

func check_string_alphabet(str string) bool {  
	for _, charVariable := range str {  
	 if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {  
	  return false  
	 }  
	}  
	return true  
   }  


func main()  {
	fmt.Println(check_string_alphabet("Amber"))
	fmt.Println(check_string_alphabet("1234"))
}

// Regex
// ^ - Beginning of a string
// [ - Character group starting
// a-z - lower case letter
// A-Z - upper case letter
// ] - Character group end
// $ - String ending
// + - one or more characters