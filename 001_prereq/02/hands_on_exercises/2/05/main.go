package main

import (
	"fmt"
)

func main() {
	type person struct {
		fName   string
		lName   string
		favFood []string
	}

	p1 := person{"James", "Bond", []string{"Steak", "Chips", "Martini"}}
	fmt.Println(p1.favFood)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}

}
