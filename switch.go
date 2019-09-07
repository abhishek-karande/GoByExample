package main

import "fmt"
import "time"

func switchfunc() {
	i := 0
	mod := i % 2
	switch mod {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	}
	fmt.Println(time.Now().Weekday())
}
