package main

import "fmt"

func slicesFunc() {
	fmt.Println("Executing slicesFunc")
	s := make([]string, 3)
	s[0] = "new"
	s[1] = "slice"
	s[2] = "created"
	s = append(s, "extra ")
	s = append(s, "elements")
	s = append(s, "added")
	fmt.Println("printing slices:", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("printing copied slice", c)
}
