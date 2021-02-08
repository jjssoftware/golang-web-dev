package main

import "fmt"

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	s = string(bs)
	fmt.Println(s)

	fmt.Println(string(bs[:14]))
	fmt.Println(string(bs[10:22]))
	fmt.Println(string(bs[17:]))

	for _, v := range bs {
		fmt.Printf("%#U\n", v)
	}
}
