package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{fourWheel: true, vehicle: vehicle{doors: 2, colour: "green"}}

	s := sedan{luxury: true, vehicle: vehicle{doors: 4, colour: "blue"}}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.fourWheel)
	fmt.Println(s.luxury)
}
