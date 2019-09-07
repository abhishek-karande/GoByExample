package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 66

	return &p
}

func structsDemo() {
	fmt.Println("Executing strucs demo")
	fmt.Println(person{"abhishek", 23})
	fmt.Println(person{name: "mark"})
	fmt.Println(&person{name: "cuban", age: 50})
	fmt.Println(*newPerson("Daymond"))

	rob := person{name: "robert", age: 48}
	fmt.Println(rob.name)

	rob2 := &rob
	fmt.Println(rob2.name)

	rob2.name = "Baratheon" //since rob2 is refering to rob ,rob.name will also get changed
	fmt.Println(rob2.name)
	fmt.Println(rob.name)
}
