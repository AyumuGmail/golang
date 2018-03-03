package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	w2 := Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w2)

	w.Circle.Center.X = 42
	//w.X = 42
	//w2.X = 24
	w2.Circle.Center.X = 24

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w2)

}
