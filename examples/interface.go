// An Interface is an abstract type.

// Interface describes all the methods of a
//  method set and provides the signatures for each method.

// To create interface use interface keyword, 
// followed by curly braces containing a list of 
// method names, along with any parameters or return 
// values the methods are expected to have.

// An interfaces act as a blueprint for method sets, 
// they must be implemented before being used. Type 
// that satisfies an interface is said to implement it.

package main

import (
	"fmt"
)

type Employee interface{
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}
type Emp int
func (e Emp) PrintName(name string)  {
	fmt.Println("Employee id", e)
	fmt.Println("Employee Name:\t", name)
}

func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}