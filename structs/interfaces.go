package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Squares struct {
	Length float64
}

func (c Squares) Area() float64 {
	return c.Length * c.Length
}

func (c Squares) Perimeter() float64 {
	return 2 * c.Length
}
