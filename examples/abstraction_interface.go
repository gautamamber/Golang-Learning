package main

import "fmt"

// Define a new data type triangle
type Triangle struct{
	base, height float32
}

// Define a new  data type square
type Square struct{
	length float32
}

// Define a new data type Rectangle
type Rectangle struct{
	length, width float32
}

// Define a new  data type Circle
type Circle struct{
	radius float32
}

type Area interface{
	Area() float32 // Return float value
}

// Triangle area method

func (t Triangle) Area() float32 {
	return 0.5 * t.height * t.base
}

// Square area method

func (s Square) Area() float32  {
	return s.length * s.length
}

// Rectangle area method

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// Circle area method

func (c Circle) Area() float32  {
	return 3.14 * (c.radius * c.radius)
}

func main()  {
	t := Triangle{base: 15, height: 25}
	s := Square{length: 5}
	r := Rectangle{length: 5, width: 10}
	c := Circle{radius: 5}

	var a Area
	a = t
	fmt.Println("Area of Triangle", a.Area())
 
	// Assign to the interface a variable of type "Square"
	a = s
	fmt.Println("Area of Square", a.Area())
 
	// Assign to the interface a variable of type "Rectangle"
	a = r
	fmt.Println("Area of Rectangle", a.Area())
 
	// Assign to the interface a variable of type "Circle"
	a = c
	fmt.Println("Area of Circle", a.Area())
}

