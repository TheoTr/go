package triangle

type Triangle struct {
	side1  float64
	side2  float64
	side3  float64
	height float64
	base   float64
}

// Perimeter calculates the perimeter of givens width and height of a triangle.
func (t Triangle) Perimeter() float64 {
	return (t.side1 + t.side2 + t.side3)
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

// Area calculates the area of givens width and height of a triangle.
func (t Triangle) Area() float64 {
	return ((t.base * t.height) / 2)
}

func Area(s Shape) float64 {
	return s.Area()
}
