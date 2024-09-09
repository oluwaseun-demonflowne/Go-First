package reciever

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) printName() {
	fmt.Printf("%s is his name", p.name)
}

func Reciever() {
	person := Person{
		name: "Alex",
		age:10,
	}
	person.printName()
}