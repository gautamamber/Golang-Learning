package main

import "fmt"

var light_switch bool = false

func main()  {
	print_light_switch_state()
	toogle_switch()
	print_light_switch_state()
	toogle_switch()
}

func print_light_switch_state()  {
	fmt.Println("The light switch is off", light_switch)
}

func toogle_switch()  {
	light_switch =! light_switch
}
