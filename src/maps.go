// Kind of Hash and Dictionary

package main

import (
	"fmt"
	"sort"
)

func main()  {
	nation_capital()
}

func nation_capital()  {
	var nation map[string]string = make(map[string]string)
	nation["India"] = "Delhi"
	nation["Canada"] = "Ontario"
	nation["Afghanistan"] = "Kabul"
	nation["PAK"] = "lahore"

	fmt.Println("Print in unsorted order")
	print_nation_capital(nation)
	fmt.Println("==========")
	print_sorted_nation_capital(nation)
}

func print_sorted_nation_capital(captital map[string]string)  {
	keys:= make([]string, len(captital))
	for key, _ := range captital{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, v := range keys{
		if v == ""{
			continue
		}
		fmt.Println(captital[v])
	}
}

func print_nation_capital(captital map[string]string) {
	for key, value := range captital{
		fmt.Println("Capital is", key, value)
	}
}