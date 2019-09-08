package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 999 {
		return -1, errors.New("can't work with 999")
	} else {
		return arg * 2, nil
	}
}

type myError struct {
	arg  int
	prob string
}

func (e myError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 999 {
		return -1, myError{arg, "can't work with 999"}
	} else {
		return arg * 2, nil
	}
}
func ErrorsDemo() {
	fmt.Println(f1(999))
	fmt.Println(f1(4))
	fmt.Println(f2(999))
	fmt.Println(f2(4))

}
