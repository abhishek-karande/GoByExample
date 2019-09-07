package main

import "fmt"

func rangeFunc() {
	fmt.Println("executing rangeFunc")
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums { //first value "_" actually gives the index
		sum += num
	}
	m := make(map[string]int)
	m["salary"] = 1000
	m["exp"] = 5
	for k, v := range m {
		fmt.Printf("%s -> %d ", k, v)
		fmt.Println()
	}
	for k := range m {
		fmt.Println(k, "->", m[k])
	}
}
