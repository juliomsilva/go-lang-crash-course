package main

import (
	"fmt"
	"strconv"
)

type person struct {
	firstName string
	lastName  string
	gender    string
	age       int
}

func (p person) greeting() string {
	return "Hi, I am " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "."
}

func (p *person) hasBirthday() {
	p.age++
}

func main() {
	samantha := person{firstName: "Samantha", lastName: "Smith", gender: "f", age: 25}
	robert := person{"Robert", "Martin", "m", 52}

	fmt.Println(samantha)
	fmt.Println(robert)

	samantha.hasBirthday()
	fmt.Println(samantha.greeting())
}
