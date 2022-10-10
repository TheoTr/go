package geom

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

func Area(s Shape) float64 {
	return s.Area()
}

func PerimeterRectangle(rect Rectangle) float64 {
	return 20
}

// Rectangle
type Rectangle struct {
	height float64
	width  float64
}

// Perimeter calculates the perimeter of a given height and width of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

// Area calculates the area of a given height and weight of a rectangle.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle
type Circle struct {
	radius float64
}

// Perimeter calculates the perimeter of a given radius of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area calculates the area of a given radius of a circle.
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// Pentagone
type Pentagone struct {
	edge     float64
	apotheme float64
}

// Perimeter calculates the perimeter of a given edge of a pentagone.
func (p Pentagone) Perimeter() float64 {
	return 5 * p.edge
}

// Area calculates the area of a given apotheme of a pentagone.
func (p Pentagone) Area() float64 {
	return 0.5 * p.Perimeter() * p.apotheme
}

// Hexagone
type Hexagone struct {
	width float64
}

// Perimeter calculates the perimeter of a given width of a hexagone.
func (h Hexagone) Area() float64 {
	return (3 * math.Sqrt(3) * h.width * h.width) / 2
}

// Area calculates the area of a given width of a hexagone.
func (h Hexagone) Perimeter() float64 {
	return h.width * 6
}

// Triangle
type Triangle struct {
	side1  float64
	side2  float64
	side3  float64
	height float64
	base   float64
}

// Perimeter calculates the perimeter of a given side1, side2 and side3 of a triangle.
func (t Triangle) Perimeter() float64 {
	return (t.side1 + t.side2 + t.side3)
}

// Area calculates the area of a given base and height of a triangle.
func (t Triangle) Area() float64 {
	return ((t.base * t.height) / 2)
}

type Star struct {
	length float64
}

func (s Star) Perimeter() float64 {
	return (s.length + s.length) * 5
}

func (s Star) Area() float64 {
	return (s.length * s.length) * 10
}
