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

func (t truck) transportationDevice() string {
	return fmt.Sprint("A ", t.doors, " door truck moves heavy goods. Is four wheel: ", t.fourWheel)
}

func (s sedan) transportationDevice() string {
	return fmt.Sprint("A ", s.colour, " sedan moves people in style. Is luxury: ", s.luxury)
}

func main() {
	t := truck{fourWheel: true, vehicle: vehicle{doors: 2, colour: "green"}}

	s := sedan{luxury: true, vehicle: vehicle{doors: 4, colour: "blue"}}

	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())
}
