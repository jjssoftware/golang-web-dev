package main

import "fmt"

func main() {
	m := map[string]int{
		"Joe":  49,
		"Syl":  50,
		"Jess": 19,
	}

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
