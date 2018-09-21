package main

import(
 "fmt"
 "math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{width: 5, height: 10}
	c := circle{radius: 12.4}
	measure(r)
	measure(c)
}
