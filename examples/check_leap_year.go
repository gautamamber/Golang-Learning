package main

import (
	"fmt"
)

func isLeapYear(year int) bool  {
	leap_year := false
	if year % 4 == 0{
		if year % 100 == 0{
			if year % 400 == 0{
				leap_year = true
			} else {
				leap_year = false
			}
		} else {
			leap_year = false
		}
	} else {
		leap_year = false
	}
	return leap_year
}

func main()  {
	leap_year := isLeapYear(2000)
	bool1 := isLeapYear(1997)
	fmt.Println(leap_year)
	fmt.Println(bool1)
}