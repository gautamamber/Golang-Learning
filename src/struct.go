package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func new_person(name string) *person  {
	p := person{name:name}
	p.age = 42
	return &p
}

func main()  {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(new_person("John"))
}