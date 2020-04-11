package main  
  
import (  
 "fmt"  
 "strconv"  
)  
  
func main() {  
 var octal string  
 fmt.Print("Enter Octal Number:")  
 fmt.Scanln(&octal)  
 output, err := strconv.ParseInt(octal, 8, 64)  
 if err != nil {  
  fmt.Println(err)  
  return  
 }  
  
 fmt.Printf("Output %d", output)  
}  