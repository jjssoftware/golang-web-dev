package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func main() {
	var g1 gator
	g1 = 123
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	x = 456
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1) //converting g1 to int works
	fmt.Println(x)

	g1.greeting()

}
