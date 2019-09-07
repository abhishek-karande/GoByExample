package main

import "fmt"

func mapsFunc() {
	fmt.Println("Executing mapsFunc")
	m := make(map[string]int)
	m["salary"] = 1000
	m["exp"] = 5
	fmt.Println("map-->", m)
	delete(m, "exp")
	_, prs := m["exp"] //upon accessing a key it return two values 1. value of key 2. true/ false showing presence of key

	fmt.Println(prs)

}
