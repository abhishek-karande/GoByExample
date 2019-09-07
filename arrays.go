package main

import "fmt"

func arrayFunc() {
	fmt.Println("Performing array ops")
	var a [5]int
	fmt.Println("array", a)
	a[4] = 100
	fmt.Println("array", a)
	fmt.Println("element at 4:", a[4])
	fmt.Println("length", len(a))

	//array with va;ues assigned at the time of delcaration
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("pre initialized array ", b)
	//2D array
	var twoDim [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("Two dim array", twoDim)
	fmt.Println(len(twoDim))
}
