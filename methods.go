package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.height * r.width
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height

}

func methodDemo() {
	r := rect{10, 20}

	fmt.Println(r.area())
	fmt.Println(r.perim())

}
