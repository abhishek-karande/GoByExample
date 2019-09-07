package main

import "fmt"

func closureFunc() func() int {
	fmt.Println("Calling closure...")
	fmt.Println("Closure is nothing but an anonymous function")
	i := 0
	return func() int {
		i++
		return i
	}
}
