package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of the given dimensions
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the area of the given dimensions
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
