package main

import "fmt"

func main() {
	type person struct {
		fName string
		lName string
	}

	p1 := person{"James", "Bond"}
	fmt.Println(p1)
	fmt.Println(p1.fName)
}
