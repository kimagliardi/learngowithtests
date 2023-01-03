package structs

import "math"

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Triangle struct {
	Base   float64
	Height float64
}

// In Go interface resolution is implicit.
// If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)

}
