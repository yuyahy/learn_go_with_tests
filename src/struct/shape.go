package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)

}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
