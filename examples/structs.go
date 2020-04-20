package main

import "fmt"

type rectangle struct{
	length float64
	breadth float64
	color string
}

type rectangle_two struct{
	length int
	breadth int
	color string
	geometry struct{
		area int
		perimeter int
	}
}

func main()  {
	fmt.Println(rectangle{10.4, 20.3, "red"})
	var rect rectangle_two
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"
	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)
	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
	rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)
}

// An instance of a struct can also be created with 
// the new keyword. It is then possible to assign data 
// values to the data fields using dot notation


// Method of assigning a custom default value can be 
// achieve by using constructor function. Instead of 
// creating a struct directly, the Info function can 
// be used to create an Employee struct with a custom
//  default value for the Name and Age field.

// func (obj *Employee) Info() {
// 	if obj.Name == "" {
// 		obj.Name = "John Doe"
// 	}
// 	if obj.Age == 0 {
// 		obj.Age = 25
// 	}
// }

// Above is constructor

// The reflect package support to check the underlying type of a struct.

// var rect1 = rectangle{10, 20, "Green"}
// 	fmt.Println(reflect.TypeOf(rect1)) 