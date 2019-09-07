package main

import "fmt"

func main() {

	fmt.Println("Hello Go!...again")
	//switchfunc()
	//arrayFunc()
	//slicesFunc()
	//mapsFunc()
	//rangeFunc()
	/*sum, mul := multiRetFunc(5, 6)
	fmt.Println("sum-->", sum, " mul--> ", mul)*/

	/*sum := variFunc(3, 4, 5, 6)
	fmt.Println(sum)*/

	/*i := closureFunc()
	fmt.Println(i())*/

	// fact := fact(9)
	// fmt.Println(fact)

	// pointerDemo()

	//structsDemo()
	//methodDemo()

	r := rect1{4, 5}
	c := circle{7.69}
	measure(r)
	measure(c)
}
