package example_interfaces

import (
	"fmt"
	"math"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) Geometry {
	switch geoType {
	case rectType:
		return Rect{width: values[0], height: values[1]}
	case circleType:
		return Circle{radius: values[0]}
	}
	return nil
}

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2 * r.width * 2 * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func details(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Init() {
	r := newGeometry(rectType, 2, 4)
	details(r)
	c := newGeometry(circleType, 5)
	details(c)
}
