package main  
import (  
 	"fmt"  
)  
  
func main() {  
 //reading an string  
 var name string  
 fmt.Print("What is your Name? ")  
 fmt.Scan(&name)  
 fmt.Println("Entered Name is :", name)  
 multiple()
}  

func multiple() {  
	//reading an multiple input string  
	var fname string  
	var lname string  
	fmt.Print("What is your Name? ")  
	fmt.Scanln(&fname, &lname)  
	fmt.Printf("Entered Name is %s %s", fname, lname)  
   }  
   