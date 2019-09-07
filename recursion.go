package main

import "fmt"

func fact(n int) int {
	fmt.Println("Calculating factorial with recursion ")
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
