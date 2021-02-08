package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 123
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	x = 456
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = g1 //does not work, different types

}
