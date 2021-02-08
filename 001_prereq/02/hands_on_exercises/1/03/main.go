package main

import "fmt"

type person struct {
	fname string
	lname string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) Speak() {
	fmt.Println(p.fname, p.lname, "is a person who says Hi")
}

func (sa secretAgent) Speak() {
	fmt.Println(sa.fname, sa.lname, "is a secret agent with license to kill", sa.ltk)
}

type human interface {
	Speak()
}

func saySomething(h human) {
	h.Speak()
}

func main() {
	p := person{"Joe", "Lippa", 49}
	sa := secretAgent{
		person{
			"James",
			"Bond",
			35,
		},
		true,
	}
	saySomething(p)
	saySomething(sa)
}
