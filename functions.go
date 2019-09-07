package main

import "fmt"

func multiRetFunc(a, b int) (int, int) {
	fmt.Println("Executing multiple return values function...")
	sum := a + b
	mul := a * b
	return sum, mul
}
func variFunc(nums ...int) int {
	fmt.Println("Executing variadic functions")
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
