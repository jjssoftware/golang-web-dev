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

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, "is a person who says Hi")
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fname, sa.lname, "is a secret agent with license to kill", sa.ltk)
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
	fmt.Println(p.fname)
	p.pSpeak()

	fmt.Println(sa.fname)
	sa.saSpeak()
	sa.pSpeak()
}
