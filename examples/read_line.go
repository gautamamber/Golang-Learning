package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	file, err := os.Open("open.txt")
	if err != nil{
		log.Fatalf("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	
	for scanner.Scan(){
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	for _, eachline := range txtlines{
		fmt.Println(eachline)
	}
}