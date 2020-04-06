package main

import (
	"fmt"
)

func main()  {
	// multidimensional matrix
	my_matrix := [3][4]int{{1,2,3,4}, {2,4,3,4}, {4,3,2,2}}
	fmt.Println(my_matrix[0][0])
	fmt.Println(my_matrix[1][0])
	fmt.Println(my_matrix[2][0])
	fmt.Println("\n\n")
	fmt.Println(my_matrix)

	print_matrix(my_matrix)
}

func print_matrix(input [3][4]int) {
	row_length := len(input)
	column_length := len(input[0])
	for i := 0; i<row_length; i++{
		for j:=0; j< column_length; j++{
			fmt.Printf("%5d ", input[i][j])
		}
		fmt.Println()
	}
}