package main

import (
	"fmt"
)

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, p.lName, "is walking")
}

func main() {
	p1 := person{"James", "Bond", []string{"Steak", "Chips", "Martini"}}
	fmt.Println(p1.walk())
}