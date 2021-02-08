package main

import "fmt"

func main() {
	xi := []int{1, 2, 3}
	fmt.Println(xi)

	for i := range xi {
		fmt.Println("idx:", i)
	}

	for i2, v := range xi {
		fmt.Println("idx:", i2, ", value:", v)
	}
}
