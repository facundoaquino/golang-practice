package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Color() string
}

type Circle struct {
	Radius float64
}

type Rectangule struct {
	Width  float64
	Height float64
	ColorH string
}

func (c *Circle) Area() float64 {
	return 3.1416 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.1416 * c.Radius
}

func printShape(s Shape) {
	fmt.Println("First show the color: ", s.Color())
	fmt.Println(s.Area())
}

type Triangule struct {
	Base   float64
	Height float64
	ColorH string
}

func (t *Triangule) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t *Triangule) Perimeter() float64 {
	return t.Base + t.Height + (math.Sqrt(t.Base*t.Base + t.Height*t.Height))
}

func (t *Triangule) Color() string {
	return "Color: " + t.ColorH

}

// EL RECTANGULO TIENE ESTOS METODOS IMPLEMENTADOS, ENTONCES ==> ES UNA SHAPE

func (r Rectangule) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangule) Perimeter() float64 {
	return 2 * r.Width * r.Height
}

func (r *Rectangule) Color() string {
	return "Color: " + r.ColorH
}

func main() {

	rec := Rectangule{
		Width:  10,
		Height: 20,
		ColorH: "aad158",
	}

	printShape(&rec)

	printShape(&Triangule{Base: 10, Height: 20, ColorH: "aad159"})
}
