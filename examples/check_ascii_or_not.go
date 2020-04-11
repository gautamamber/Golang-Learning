package main

import (
	"fmt"
)
func main() {  
	charVariable := 'T' // This is infered as rune data type  
	asciiCode := int(charVariable)  
	fmt.Printf("%s ASCII code is : %d", string(charVariable), asciiCode)  
   }  