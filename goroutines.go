package main

import "fmt"

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func goRoutinesDemo() {
	f("Synchronous call")
	go f("goroutine")

	go func(str string) {
		fmt.Println(str)
	}("anonymous")
	go f("another goroutine")

	fmt.Scanln()
	fmt.Println("Ciao")
}
