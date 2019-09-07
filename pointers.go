package main

import "fmt"

func changeValFunc(ival int) {
	fmt.Println("pass by value...")
	fmt.Println(&ival)

	ival = 0
}

func changeValPtrFunc(iptr *int) {
	fmt.Println("pass by reference")
	fmt.Println(iptr)
	*iptr = 0
}
func pointerDemo() {
	val := 100
	changeValFunc(val)
	fmt.Println(&val)
	fmt.Println("value after changeValFunc", val)
	changeValPtrFunc(&val)
	fmt.Println("value after changeValPtrFunc", val)

}
